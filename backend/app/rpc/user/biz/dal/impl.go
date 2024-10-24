package dal

import (
	"backend/app/common/constant"
	"backend/app/rpc/user/biz/model"
	"backend/library/metric"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	cache "github.com/mgtv-tech/jetcache-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

var _ UserDal = (*UserDalImpl)(nil)

type UserDalImpl struct {
	cache                      cache.Cache
	db                         *gorm.DB
	userByIdCache              *cache.T[int64, *model.User]
	userRelevantCountByIdCache *cache.T[int64, *model.UserRelevantCount]
	redis                      redis.UniversalClient
}

func (s *UserDalImpl) AddTokenToBlackList(ctx context.Context, token string) error {
	key := fmt.Sprintf("%s:%s", constant.UserTokenBlackListKey, token)
	_, err := s.redis.Set(ctx, key, 1, time.Hour*24).Result()
	if err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromRedisUserTokenBlackList)
		klog.Errorf("redis.Set(%s) failed, err:%v", key, err)
		return err
	}
	return nil
}

func (s *UserDalImpl) ExistUserByID(ctx context.Context, id int64) (bool, error) {
	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *UserDalImpl) GetUserRelevantCountByID(ctx context.Context, id int64) (*model.UserRelevantCount, error) {
	count := &model.UserRelevantCount{}
	var err error
	err = s.userRelevantCountByIdCache.Once(
		ctx, userRelevantCountFromIdKey(id),
		cache.Value(count), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			var err error
			cnt := &model.UserRelevantCount{}
			if cnt, err = getUserRelevantCountByID(ctx, s.db, id); err == nil {
				return cnt, nil
			} else {
				metric.IncrGauge(metric.LibClient, constant.PromDBRelevantCount)
				klog.Errorf("db.getUserRelevantCountByID(%d) failed, err:%v", id, err)
				return nil, err
			}
		},
		),
	)
	return count, err
}

func (s *UserDalImpl) ExistUserByUserName(ctx context.Context, name string) (bool, error) {
	id, err := getCacheUsernameToUid(s.cache, ctx, name)
	if err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
		klog.Errorf("cache.getCacheUsernameToUid(%s) failed, err:%v", name, err)
	}
	if id > 0 {
		return true, nil
	}

	if f, err := existUserByUserName(ctx, s.db, name); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBUser)
		klog.Errorf("db.existUserByUserName(%s) failed, err:%v", name, err)
		return false, err
	} else {
		return f, nil
	}
}

func NewUserDalImpl(c cache.Cache, db *gorm.DB, r redis.UniversalClient) *UserDalImpl {
	return &UserDalImpl{
		cache:                      c,
		db:                         db,
		userByIdCache:              cache.NewT[int64, *model.User](c),
		redis:                      r,
		userRelevantCountByIdCache: cache.NewT[int64, *model.UserRelevantCount](c),
	}
}

func (s *UserDalImpl) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	// 数据库不存在 返回gorm.ErrRecordNotFound
	user := &model.User{}
	err := s.userByIdCache.Once(
		ctx, userFromUidKey(id),
		cache.Value(user), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			if user, err := getUserByID(ctx, s.db, id); err == nil {
				e := setCacheUsernameToUid(s.cache, ctx, user.Username, user.ID)
				if e != nil {
					metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
					klog.Errorf("cache.setCacheUsernameToUid(%s, %d) failed, err:%v", user.Username, user.ID, e)
				}
				return user, nil
			} else {
				metric.IncrGauge(metric.LibClient, constant.PromDBUser)
				klog.Errorf("db.getUserByID(%d) failed, err:%v", id, err)
				return nil, err
			}
		},
		),
	)

	return user, err
}

func (s *UserDalImpl) UpdateUserByID(ctx context.Context, id int64, user *model.User) error {

	cacheKey := userFromUidKey(id)

	if err := delCacheUsernameToUid(s.cache, ctx, user.Username); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
		klog.Errorf("cache.delCacheUsernameToUid(%s) failed, err:%v", user.Username, err)
		return err
	}

	if err := s.userByIdCache.Delete(ctx, cacheKey); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUID)
		klog.Errorf("cache.delete(%s) failed, err:%v", cacheKey, err)
		return err
	}

	if err := updateUserByID(ctx, s.db, id, user); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBUser)
		klog.Errorf("db.updateUserByID(%d) failed, err:%v", id, err)
		return err
	}

	go s.delayedDoubleDelete(cacheKey, 10, 200*time.Millisecond)
	return nil
}

func (s *UserDalImpl) DeleteUserByID(ctx context.Context, id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%d", constant.UserDetailCacheFromUidKey, id)

	var user *model.User
	if user, err = s.GetUserByID(ctx, id); err != nil {
		klog.Errorf("UserDalImpl.GetUserByID(%d) failed,err :%+v", id, err)
		return err
	}

	if err = delCacheUsernameToUid(s.cache, ctx, user.Username); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
		klog.Errorf("cache.delCacheUsernameToUid(%s) failed, err:%v", user.Username, err)
		return err
	}

	if err = s.userByIdCache.Delete(ctx, cacheKey); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUID)
		klog.Errorf("cache.delete(%s) failed, err:%v", cacheKey, err)
		return err
	}

	if err = s.db.Delete(&model.User{}, id).Error; err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBUser)
		klog.Errorf("db.delete(%d) failed, err:%v", id, err)
		return err
	}

	go s.delayedDoubleDelete(cacheKey, 10, 200*time.Millisecond)
	return nil
}

func (s *UserDalImpl) CreateUser(ctx context.Context, user *model.User) (err error) {
	if err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromDBUser)
			klog.Errorf("db.create(%+v) failed, err:%v", user, err)
			return err
		}
		userRelevantCount := &model.UserRelevantCount{
			FollowerCount:  0,
			FollowingCount: 0,
			LikeCount:      0,
			StarCount:      0,
			SelfStarCount:  0,
			SelfLikeCount:  0,
			LiveCount:      0,
			WorkCount:      0,
			FriendCount:    0,
			UserID:         user.ID,
		}
		if err := tx.Create(userRelevantCount).Error; err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromDBRelevantCount)
			klog.Errorf("db.create(%+v) failed, err:%v", userRelevantCount, err)
			return err
		}
		return nil
	}); err != nil {
		klog.Errorf("db.transaction failed, err:%v", err)
		return err
	}

	// 设置缓存
	go func() {
		if err = s.userByIdCache.Set(ctx, constant.UserDetailCacheFromUidKey, user.ID, user); err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUID)
			klog.Errorf("cache.set(%d) failed, err:%v", user.ID, err)
		}

		if err = setCacheUsernameToUid(s.cache, ctx, user.Username, user.ID); err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
			klog.Errorf("cache.setCacheUsernameToUid(%s, %d) failed, err:%v", user.Username, user.ID, err)
		}
	}()

	return nil
}

func (s *UserDalImpl) GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]*model.User, error) {
	user := s.userByIdCache.MGet(ctx, constant.UserDetailCacheFromUidKey, ids, func(ctx context.Context, ids []int64) (map[int64]*model.User, error) {
		var users []model.User
		err := s.db.Find(&users, "id in (?)", ids).Error
		if err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromDBUser)
			klog.Errorf("db.find(%v) failed, err:%v", ids, err)
			return nil, err
		}

		userMap := make(map[int64]*model.User, len(users))
		for _, user := range users {
			userMap[user.ID] = &user
		}
		return userMap, nil
	})
	return user, nil
}

func (s *UserDalImpl) GetFollowersByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error) {

	//err := s.db.WithContext(ctx).Preload("Followers", func(db *gorm.DB) *gorm.DB {
	//	return db.Select("users.id, users.username, users.nickname").Limit(total).Offset(offset)
	//}).First(&model.User{}, "id = ?", userID).Error
	var followerIDs []int64

	err := s.db.WithContext(ctx).Raw(`
        SELECT follower_id 
        FROM user_follows 
        WHERE followed_id = ? 
        LIMIT ? OFFSET ?`, userID, total, offset).Scan(&followerIDs).Error

	if err != nil {
		return nil, err
	}

	ret := s.userByIdCache.MGet(ctx, constant.UserDetailCacheFromUidKey, followerIDs, func(ctx context.Context, ids []int64) (map[int64]*model.User, error) {
		return s.GetUsersByIDs(ctx, ids)
	})
	return ret, nil
}

func (s *UserDalImpl) GetFollowingsByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error) {
	var followingIDs []int64

	err := s.db.WithContext(ctx).Raw(`
        SELECT followed_id 
        FROM user_follows 
        WHERE follower_id = ? 
        LIMIT ? OFFSET ?`, userID, total, offset).Scan(&followingIDs).Error

	if err != nil {
		return nil, err
	}

	ret := s.userByIdCache.MGet(ctx, constant.UserDetailCacheFromUidKey, followingIDs, func(ctx context.Context, ids []int64) (map[int64]*model.User, error) {
		return s.GetUsersByIDs(ctx, ids)
	})

	return ret, nil
}

func (s *UserDalImpl) GetFriendsByUserID(ctx context.Context, userID int64, total, offset int) (map[int64]*model.User, error) {
	var followingIDs []int64

	// 第一步：获取用户关注的人的ID
	err := s.db.WithContext(ctx).Raw(`
        SELECT followed_id 
        FROM user_follows 
        WHERE follower_id = ?
        LIMIT ? OFFSET ?`, userID, total, offset).Scan(&followingIDs).Error

	if err != nil {
		return nil, err
	}

	if len(followingIDs) == 0 {
		return make(map[int64]*model.User), nil // 如果没有关注的人，直接返回空结果
	}

	var friendIDs []int64

	// 第二步：从关注的用户中查找哪些用户也关注了当前用户
	err = s.db.WithContext(ctx).Raw(`
        SELECT follower_id 
        FROM user_follows 
        WHERE followed_id = ? AND follower_id IN (?)`, userID, followingIDs).Scan(&friendIDs).Error

	if err != nil {
		return nil, err
	}

	// 使用缓存机制获取用户详细信息
	ret := s.userByIdCache.MGet(ctx, constant.UserDetailCacheFromUidKey, friendIDs, func(ctx context.Context, ids []int64) (map[int64]*model.User, error) {
		return s.GetUsersByIDs(ctx, ids)
	})

	return ret, nil
}

func (s *UserDalImpl) GetUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	var user *model.User
	var (
		id  int64
		err error
	)

	if id, err = getCacheUsernameToUid(s.cache, ctx, userName); err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
		klog.Errorf("cache.getCacheUsernameToUid(%s) failed, err:%v", userName, err)
	}

	if id > 0 {
		user, err = s.GetUserByID(ctx, id)
		if err != nil {
			klog.Errorf("UserDalImpl.GetUserByID(%d) failed,err :%+v", id, err)
			return nil, err
		}
		return user, nil
	} else {
		if user, err = getUserByUserName(ctx, s.db, userName); err != nil {
			metric.IncrGauge(metric.LibClient, constant.PromDBUser)
			klog.Errorf("db.getUserByUserName(%s) failed, err:%v", userName, err)
			return nil, err
		}
		go func() {
			if err = setCacheUsernameToUid(s.cache, ctx, user.Username, user.ID); err != nil {
				metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUsername)
				klog.Errorf("cache.setCacheUsernameToUid(%s, %d) failed, err:%v", user.Username, user.ID, err)
			}
			if err = s.userByIdCache.Set(ctx, constant.UserDetailCacheFromUidKey, user.ID, user); err != nil {
				metric.IncrGauge(metric.LibClient, constant.PromCacheUserDetailUID)
				klog.Errorf("cache.userByIdCache.set(%d) failed, err:%v", user.ID, err)
			}
		}()

		return user, nil
	}

}

func (s *UserDalImpl) GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := s.db.WithContext(ctx).First(user, "phone = ?", phone).Error
	if err != nil {
		return nil, err
	}

	_ = s.userByIdCache.Set(ctx, constant.UserDetailCacheFromUidKey, user.ID, &user)
	return &user, nil
}

func (s *UserDalImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := s.db.WithContext(ctx).First(user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	_ = s.userByIdCache.Set(ctx, constant.UserDetailCacheFromUidKey, user.ID, &user)
	return &user, nil
}

func (s *UserDalImpl) delayedDoubleDelete(cacheKey string, maxRetries int, initialDelay time.Duration) {
	delay := initialDelay
	for attempt := 1; attempt <= maxRetries; attempt++ {
		time.Sleep(delay)
		err := s.cache.Delete(context.Background(), cacheKey)
		if err == nil {
			klog.Infof("Successfully performed delayed delete for cache key %s on attempt %d", cacheKey, attempt)
			return
		}

		klog.Errorf("Attempt %d: Failed to delete cache for key %s: %v", attempt, cacheKey, err)

		delay *= 2
	}
	klog.Errorf("Exceeded maximum retries (%d) for deleting cache key %s", maxRetries, cacheKey)
}

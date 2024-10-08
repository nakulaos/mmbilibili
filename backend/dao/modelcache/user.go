package modelcache

import (
	"backend/common/constant"
	"backend/dao/model"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"time"
)

func SetUserModelCache(cacheImpl *cache.Cache, newUser *model.User) {
	_ = cacheImpl.Set(&cache.Item{
		Key:   constant.UserInfoCacheUidKey + fmt.Sprintf("%d", newUser.ID),
		Value: newUser,
		TTL:   time.Minute * 30,
	})
	_ = cacheImpl.Set(&cache.Item{
		Key:   constant.UserNameToUidKey + newUser.Username,
		Value: newUser.ID,
		TTL:   time.Minute * 30,
	})
}

func GetUserModelCacheFromId(cacheImpl *cache.Cache, ctx context.Context, userId uint) (userModel *model.User, err error) {
	userModel = &model.User{}
	_ = cacheImpl.Get(ctx, constant.UserInfoCacheUidKey+fmt.Sprintf("%d", userId), userModel)
	return
}

func GetUserModelCacheFromUsername(cacheImpl *cache.Cache, ctx context.Context, username string) (userModel *model.User, err error) {
	userModel = &model.User{}
	var userId uint
	err = cacheImpl.Get(ctx, constant.UserNameToUidKey+username, &userId)
	if err != nil || userId == 0 {
		return userModel, nil
	}
	_ = cacheImpl.Get(ctx, constant.UserInfoCacheUidKey+fmt.Sprintf("%d", userId), userModel)
	return userModel, nil
}

func DelUserModelCache(cacheImpl *cache.Cache, ctx context.Context, userId uint) (err error) {
	key := constant.UserInfoCacheUidKey + fmt.Sprintf("%d", userId)
	for i := 0; i < 10; i++ {
		err = cacheImpl.Delete(ctx, key)
		if err == nil || err == cache.ErrCacheMiss {
			break
		}
	}
	return err
}

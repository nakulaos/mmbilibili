package livebusinessrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/tools"
	"backend/pkg/xerror"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
	"gorm.io/gorm"
	"time"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLiveLogic {
	return &StartLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartLiveLogic) StartLive(in *live.StartLiveReq) (*live.LiveDetailResp, error) {
	var err error
	uid := in.Uid
	streamPusherURL := fmt.Sprintf("%s/%s/live/%d/%d", l.svcCtx.Config.App.LivePusherUrl, l.svcCtx.Config.App.AppName, uid, time.Now().Unix())
	liveModel := model.Live{
		UID:         uint(uid),
		Title:       in.Title,
		Description: in.Description,
		PlayURL:     streamPusherURL,
		CoverURL:    in.Cover,
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		Partition:   in.Partition,
	}
	userModel, _ := modelcache.GetUserModelCacheFromId(l.svcCtx.Cache, l.ctx, uint(uid))
	if userModel.ID == 0 {
		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerror.UserNotExistError, "[StartLive] get user failed: %v", err)
			} else {
				return nil, errors.Wrapf(xerror.ServerError, "[StartLive] get user failed: %v", err)
			}
		}
	}
	liveModel.RoomID = userModel.RoomID
	liveModel.PushToken, liveModel.PlayToken = GenerateToken()
	if constant.HasStatus(userModel.Status, constant.UserNotAllowLive) {
		return nil, errors.Wrapf(xerror.UserNotAllowLiveError, "[StartLive] user not allow live")
	}

	for _, category := range in.Category {
		var categoryModel = &model.Category{}
		categoryModel, _ = modelcache.GetCategoryFromNameAndPartition(l.svcCtx.Cache, l.ctx, category, in.Partition)
		if categoryModel.ID == 0 {
			categoryModel, err = l.svcCtx.Dao.Category.WithContext(l.ctx).Where(l.svcCtx.Dao.Category.Name.Eq(category)).First()
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[StartLive] get category failed: %v", err)
			}
		}
		liveModel.Categories = append(liveModel.Categories, *categoryModel)
		go func() {
			modelcache.SetCategoryModelCache(l.svcCtx.Cache, in.Partition, category, categoryModel)
		}()
	}

	q := l.svcCtx.Dao.Begin()
	defer q.Commit()
	err = q.Live.WithContext(l.ctx).Create(&liveModel)
	if err != nil {
		q.Rollback()
		return nil, errors.Wrapf(xerror.ServerError, "[StartLive] create live failed: %v", err)
	}

	err = statistics.InitializeLiveRelevantCount(l.svcCtx.Redis, l.ctx, int(liveModel.ID))
	if err != nil {
		q.Rollback()
		return nil, errors.Wrapf(xerror.ServerError, "[StartLive] initialize live relevant count failed: %v", err)
	}

	go func() {
		modelcache.SetLiveModelCache(l.svcCtx.Cache, &liveModel)
		AddRecommendInfo(l.svcCtx.Redis, l.svcCtx.GorseClient, l.Logger, l.ctx, &liveModel)
	}()

	userStat, err := statistics.GetUserRelevantCount(l.svcCtx.Redis, l.ctx, int(userModel.ID))

	author := live.User{
		Id:             uint32(userModel.ID),
		Username:       userModel.Username,
		Nickname:       userModel.Nickname,
		Avatar:         userModel.Avatar,
		Gender:         uint32(userModel.Gender),
		Role:           uint32(userModel.Role),
		FollowerCount:  int32(userStat.FollowerCount),
		FollowingCount: int32(userStat.FollowingCount),
		LikeCount:      int32(userStat.LikeCount),
		StarCount:      int32(userStat.StarCount),
		SelfStarCount:  int32(userStat.SelfStarCount),
		SelfLikeCount:  int32(userStat.SelfLikeCount),
		LiveCount:      int32(userStat.LiveCount),
		WorkCount:      int32(userStat.WorkCount),
		FriendCount:    int32(userStat.FriendCount),
		Phone:          userModel.Phone,
		Email:          userModel.Email,
	}
	var resp live.LiveDetailResp
	resp.LiveInfo = &live.LiveInfo{
		StartTime:    liveModel.StartTime.Unix(),
		Title:        liveModel.Title,
		Description:  liveModel.Description,
		WatchCount:   0,
		LikeCount:    0,
		CommentCount: 0,
		ShareCount:   0,
		IsLike:       false,
		IsFollow:     false,
		IsStar:       false,
		IsSelf:       true,
		UserId:       uint32(userModel.ID),
		LiveId:       uint32(liveModel.ID),
		Cover:        liveModel.CoverURL,
		PlayerUrl:    liveModel.PlayURL,
		Author:       &author,
		Category:     in.Category,
		Tags:         in.Tags,
		Partition:    in.Partition,
		RoomId:       uint32(userModel.RoomID),
		Token:        liveModel.PushToken,
	}

	return &resp, nil

}

func GenerateToken() (string, string) {
	pushToken := tools.GenerateRandomToken()
	playToken := tools.GenerateRandomToken()
	return pushToken, playToken
}

func AddRecommendInfo(redisClient *redis.Redis, gorseClient *client.GorseClient, logger logx.Logger, ctx context.Context, liveModel *model.Live) {
	// redis 缓存
	//	CategoryUserPortrait string = "CategoryUserPortrait:uid:" // 分类下的用户头像 eg: CategoryUserPortrait:1 FpsGame 30
	//	HotRoomTags          string = "HotRoomTags:"              // 热门房间标签 eg: HotRoomTags:FpsGame  room_id  30
	//	UserTokenCount       string = "UserTokenCount:Uid:"       // 发弹幕所用时间 eg: UserTokenCount:Uid:1  30
	var z = make([]redis.Pair, 0)
	for _, category := range liveModel.Categories {
		z = append(z, redis.Pair{
			Score: int64(constant.StartLiveScore),
			Key:   fmt.Sprintf("%s", category.Name),
		})
		_, err := redisClient.ZaddsCtx(ctx, fmt.Sprintf("%s%s", constant.HotRoomTags, category.Name), redis.Pair{
			Score: int64(constant.InitRoomScore),
			Key:   fmt.Sprintf("%d", liveModel.RoomID),
		})
		if err != nil {
			logger.Errorf("[AddRecommendInfo] set HotRoomTags failed: %v", err)
		}
	}
	_, err := redisClient.ZaddsCtx(ctx, fmt.Sprintf("%s%d", constant.CategoryUserPortrait, liveModel.UID), z...)
	if err != nil {
		logger.Errorf("[AddRecommendInfo] set CategoryUserPortrait failed: %v", err)
	}

	// gorse
	var item client.Item
	item.ItemId = fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID)
	for _, category := range liveModel.Categories {
		item.Categories = append(item.Categories, category.Name)
	}
	item.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	item.Comment = liveModel.Description
	item.Categories = append(item.Categories, liveModel.Partition)

	_, e := gorseClient.InsertItem(context.Background(), item)
	if e != nil {
		logger.Errorf("[AddRecommendInfo] insert item failed: %v", e)
	}
}

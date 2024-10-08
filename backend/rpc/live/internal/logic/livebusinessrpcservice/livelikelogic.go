package livebusinessrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
	"time"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveLikeLogic {
	return &LiveLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveLikeLogic) LiveLike(in *live.LiveLikeReq) (*live.LiveLikeResp, error) {
	// 获取用户ID
	var err error
	uid := in.UserId
	liveId := in.LiveId

	// 获取直播
	liveModel, _ := modelcache.GetLiveModelCache(l.svcCtx.Cache, l.ctx, uint(liveId))
	if liveModel.ID == 0 {
		liveModel, err = l.svcCtx.Dao.Live.WithContext(l.ctx).Preload(l.svcCtx.Dao.Live.Categories).Where(l.svcCtx.Dao.Live.ID.Eq(uint(liveId))).First()
		if err != nil {
			return nil, errors.Wrapf(xerror.LiveNotExistError, "[LiveLike] get live error: %v", err)
		}
	}

	// 判断是否结束
	if liveModel.IsOver == 1 {
		return nil, errors.Wrapf(xerror.LiveIsOverError, "[LiveLike] live is over")
	}

	// 获取用户信息
	// 从缓存获取
	userModel, _ := modelcache.GetUserModelCacheFromId(l.svcCtx.Cache, l.ctx, uint(uid))
	if userModel.ID == 0 {
		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] get user error: %v", err)
		}
	}
	var resp = &live.LiveLikeResp{}
	// 获取直播统计数
	liveStat, err := statistics.GetLiveRelevantCount(l.svcCtx.Redis, l.ctx, int(liveId))
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] get live relevant count error: %v", err)
	}
	if int(in.Action) == constant.LikeAction {
		// like
		res, err := l.svcCtx.Dao.User.FavoriteLives.WithContext(l.ctx).Where(l.svcCtx.Dao.Live.ID.Eq(uint(liveId))).Model(userModel).Find()
		if len(res) != 0 {
			resp.LikeCount = int32(liveStat.LikeCount)
			return nil, errors.Wrapf(xerror.LiveLikeIsExistError, "[LiveLike] like live is exist")
		}
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] like live error: %v", err)
		}

		err = l.svcCtx.Dao.User.FavoriteLives.WithContext(l.ctx).Model(userModel).Append(liveModel)
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] like live error: %v", err)
		}

		err = statistics.UpdateLiveLikeCount(l.svcCtx.Redis, l.ctx, int(liveId), constant.LikeAction)
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] update live like count error: %v", err)
		}

		resp.LikeCount = int32(int(liveStat.LikeCount + 1))
	} else {
		// 取消点赞
		res, err := l.svcCtx.Dao.User.FavoriteLives.WithContext(l.ctx).Where(l.svcCtx.Dao.Live.ID.Eq(uint(liveId))).Model(userModel).Find()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveLike] favorite live error: %v", err)
		}

		if len(res) == 0 {
			return nil, errors.Wrapf(xerror.LiveLikeNotExistError, "[LiveLike] favorite live is not exist")
		}

		err = l.svcCtx.Dao.User.FavoriteLives.WithContext(l.ctx).Model(userModel).Delete(liveModel)
		if err != nil {
			return resp, errors.Wrapf(xerror.ServerError, "[LiveLike] favorite live error: %v", err)
		}

		err = statistics.UpdateLiveLikeCount(l.svcCtx.Redis, l.ctx, int(liveId), constant.UnLikeAction)
		if err != nil {
			return resp, errors.Wrapf(xerror.ServerError, "[LiveLike] update live like count error: %v", err)
		}

		resp.LikeCount = int32(liveStat.LikeCount - 1)
	}

	// 保存直播信息
	err = l.svcCtx.Dao.Live.WithContext(l.ctx).Save(liveModel)
	if err != nil {
		return resp, errors.Wrapf(xerror.ServerError, "[LiveLike] save live error: %v", err)
	}

	//go func(Cache *cache.Cache, l.ctx context.Context, key string) {
	//	err := Cache.Delete(l.ctx, key)
	//	if err != nil {
	//		l.svcCtx.Logger.Errorf("[EndLive] delete cache error: %v", err)
	//	}
	//}(l.svcCtx.Cache, l.ctx, constant.LiveInfoCacheKey+fmt.Sprintf("%d", in.LiveID))
	//go func() {
	//	// 删除缓存
	//	err = l.svcCtx.Cache.Delete(l.ctx, constant.UserInfoCacheUidKey+fmt.Sprintf("%d", uid))
	//	if err != nil {
	//		l.svcCtx.Logger.Errorf("[LiveLike] delete cache error: %v", err)
	//	}
	//	RecommendUpdateWithLikeLive(l.ctx, liveModel, in.Action)
	//}()
	go func() {
		//err = modelcache.DelUserModelCache(l.svcCtx.Cache, context.Background(), uint(uid))
		//if err != nil {
		//	l.Logger.Errorf("[LiveLike] delete cache error: %v", err)
		//}
		//err = modelcache.DelLiveModelCache(l.svcCtx.Cache, context.Background(), uint(liveId))
		//if err != nil {
		//	l.Logger.Errorf("[LiveLike] delete cache error: %v", err)
		//}

		RecommendUpdateWithLikeLive(l.svcCtx.Redis, l.svcCtx.GorseClient, l.Logger, liveModel, int(in.Action))

	}()

	return resp, nil
}

func RecommendUpdateWithLikeLive(redisClient *redis.Redis, gorseClient *client.GorseClient, logger logx.Logger, liveModel *model.Live, action int) {
	// redis
	ctx := context.Background()
	score := constant.LikeLiveScore
	if action != constant.LikeAction {
		score = -score
	}
	for _, category := range liveModel.Categories {
		_, err := redisClient.ZincrbyCtx(ctx, fmt.Sprintf("%s%s", constant.HotRoomTags, category.Name), int64(score), fmt.Sprintf("%d", liveModel.RoomID))
		if err != nil {
			logger.Errorf("[RecommendUpdateWithLikeLive] set HotRoomTags failed: %v", err)
		}
		_, err = redisClient.ZincrbyCtx(ctx, fmt.Sprintf("%s%d", constant.CategoryUserPortrait, liveModel.UID), int64(score), category.Name)
		if err != nil {
			logger.Errorf("[RecommendUpdateWithLikeLive] set CategoryUserPortrait failed: %v", err)
		}
	}

	// gorse
	if action == constant.LikeAction {
		var feedbacks []client.Feedback = make([]client.Feedback, 0)
		feedbacks = append(feedbacks, client.Feedback{
			FeedbackType: constant.LikeFeedBack,
			ItemId:       fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID),
			UserId:       fmt.Sprintf("%d", liveModel.UID),
			Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		})
		_, err := gorseClient.PutFeedback(ctx, feedbacks)
		if err != nil {
			logger.Errorf("[RecommendUpdateWithLikeLive] put feedback failed: %v", err)
		}
	} else if action == constant.UnLikeAction {
		_, err := gorseClient.DelFeedback(ctx, constant.LikeFeedBack, fmt.Sprintf("%d", liveModel.UID), fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID))
		if err != nil {
			logger.Errorf("[RecommendUpdateWithLikeLive] del feedback failed: %v", err)
		}
	}
}

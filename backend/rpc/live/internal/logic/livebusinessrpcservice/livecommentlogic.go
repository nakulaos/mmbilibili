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

type LiveCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveCommentLogic {
	return &LiveCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveCommentLogic) LiveComment(in *live.LiveCommentReq) (*live.LiveCommentResp, error) {
	// 获取用户ID
	var err error
	uid := in.UserId

	// 获取直播信息
	// 从缓存获取
	liveModel, _ := modelcache.GetLiveModelCache(l.svcCtx.Cache, l.ctx, uint(in.LiveId))
	if liveModel.ID == 0 {
		liveModel, err = l.svcCtx.Dao.Live.WithContext(l.ctx).Preload(l.svcCtx.Dao.Live.Categories).Where(l.svcCtx.Dao.Live.ID.Eq(uint(in.LiveId))).First()
		if err != nil {
			return nil, errors.Wrapf(xerror.LiveNotExistError, "[LiveCommentList] get live error: %v", err)
		}
	}

	if liveModel.IsOver == 1 {
		return nil, errors.Wrapf(xerror.LiveIsOverError, "[LiveCommentList] live is over")
	}

	commentModel := model.Danmu{
		UID:       uint(uid),
		OwnerID:   uint(in.LiveId),
		Type:      constant.DanmuLiveType,
		OwnerType: constant.Live,
		Content:   in.Content,
		SendTime:  float64(in.SendTime),
	}

	q := l.svcCtx.Dao.Begin()
	err = q.Danmu.WithContext(l.ctx).Save(&commentModel)

	if err != nil {
		l.Logger.Errorf("[LiveComment] save comment error: %v", err)
		err := q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] save comment error: %v", err)
	}

	//liveModel.CommentCount += 1
	//err = q.Live.WithContext(l.ctx).Save(liveModel)
	//if err != nil {
	//	l.Logger.Errorf("[LiveComment] save live error: %v", err)
	//	err := q.Rollback()
	//	if err != nil {
	//		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
	//	}
	//	return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] save live error: %v", err)
	//}

	err = statistics.UpdateLiveCommentCount(l.svcCtx.Redis, l.ctx, int(in.LiveId), constant.CommentAction)
	if err != nil {
		l.Logger.Errorf("[LiveComment] update live comment count error: %v", err)
		err := q.Rollback()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] rollback error: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] update live comment count error: %v", err)
	}

	err = q.Commit()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveComment] commit error: %v", err)
	}
	go func() {
		RecommendUpdateWithCommentLive(l.svcCtx.Redis, l.svcCtx.GorseClient, l.Logger, liveModel, constant.CommentAction)
		//err = modelcache.DelLiveModelCache(l.svcCtx.Cache, l.ctx, uint(in.LiveId))
		//if err != nil {
		//	l.Logger.Errorf("[LiveComment] delete cache error: %v", err)
		//}
	}()

	var resp = &live.LiveCommentResp{}
	resp.CommentId = uint32(commentModel.ID)
	return resp, nil
}

func RecommendUpdateWithCommentLive(redisClient *redis.Redis, gorseClient *client.GorseClient, logger logx.Logger, liveModel *model.Live, action int) {
	// redis
	score := constant.CommentLiveScore
	if action != constant.CommentAction {
		score = -score
	}
	ctx := context.Background()
	for _, category := range liveModel.Categories {
		_, err := redisClient.ZincrbyCtx(ctx, fmt.Sprintf("%s%s", constant.HotRoomTags, category.Name), int64(score), fmt.Sprintf("%d", liveModel.RoomID))
		if err != nil {
			logger.Errorf("[RecommendUpdateWithCommentLive] set HotRoomTags failed: %v", err)
		}
		_, err = redisClient.ZincrbyCtx(ctx, fmt.Sprintf("%s%d", constant.CategoryUserPortrait, liveModel.UID), int64(score), category.Name)
		if err != nil {
			logger.Errorf("[RecommendUpdateWithCommentLive] set CategoryUserPortrait failed: %v", err)
		}
	}

	// gorse
	if action == constant.CommentAction {
		var feedbacks []client.Feedback = make([]client.Feedback, 0)
		feedbacks = append(feedbacks, client.Feedback{
			FeedbackType: constant.CommentFeedBack,
			ItemId:       fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID),
			UserId:       fmt.Sprintf("%d", liveModel.UID),
			Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		})
		_, err := gorseClient.PutFeedback(ctx, feedbacks)
		if err != nil {
			logger.Errorf("[RecommendUpdateWithCommentLive] put feedback failed: %v", err)
		}
	} else if action == constant.UnCommentAction {
		_, err := gorseClient.DelFeedback(ctx, constant.CommentFeedBack, fmt.Sprintf("%d", liveModel.UID), fmt.Sprintf("%s:%d", constant.LiveType, liveModel.ID))
		if err != nil {
			logger.Errorf("[RecommendUpdateWithCommentLive] del feedback failed: %v", err)
		}
	}
}

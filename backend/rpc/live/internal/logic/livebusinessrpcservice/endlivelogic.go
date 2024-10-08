package livebusinessrpcservicelogic

import (
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEndLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndLiveLogic {
	return &EndLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EndLiveLogic) EndLive(in *live.EndLiveReq) (*live.LiveDetailResp, error) {
	ctx := l.ctx
	var resp = &live.LiveDetailResp{}
	liveModel, err := l.svcCtx.Dao.Live.WithContext(ctx).Where(l.svcCtx.Dao.Live.ID.Eq(uint(in.LiveId))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.LiveNotExistError, "[EndLive] get live error: %v", err)
	}

	liveModel.IsOver = 1
	liveModel.EndTime = time.Now()
	err = l.svcCtx.Dao.Live.WithContext(ctx).Save(liveModel)
	go func() {
		err := modelcache.DelLiveModelCache(l.svcCtx.Cache, ctx, liveModel.ID)
		if err != nil {
			l.Logger.Errorf("[EndLive] delete cache error: %v", err)
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[EndLive] save live error: %v", err)
	}

	userModel, err := l.svcCtx.Dao.User.WithContext(ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(in.UserId))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerror.UserNotExistError, "[EndLive] get user failed: %v", err)
		}
		return nil, errors.Wrapf(xerror.ServerError, "[EndLive] get user failed: %v", err)
	}

	userStat, err := statistics.GetUserRelevantCount(l.svcCtx.Redis, l.ctx, int(userModel.ID))
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[EndLive] get user stat error: %v", err)
	}

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
		UserId:       in.UserId,
		LiveId:       uint32(liveModel.ID),
		Cover:        liveModel.CoverURL,
		PlayerUrl:    liveModel.PlayURL,
		Author:       &author,
		EndTime:      liveModel.EndTime.Unix(),
		IsOver:       liveModel.IsOver == 1,
	}

	return resp, nil
}

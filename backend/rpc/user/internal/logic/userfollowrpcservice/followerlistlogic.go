package userfollowrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *user.FollowerListReq) (*user.FollowerListResp, error) {
	uid := in.ActionId

	if in.Total <= 0 || in.Total > int32(constant.ListDefaultLimit) {
		in.Total = int32(constant.ListDefaultLimit)
	}

	// 获取粉丝列表
	requestModel, err := l.svcCtx.Dao.User.WithContext(l.ctx).Preload(l.svcCtx.Dao.User.Followers.Limit(int(in.Total)).Offset(int(in.PageSize * in.Page))).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FollowingListLogic] get following list error: %v", err)
	}

	var follwerIds []uint
	for _, follower := range requestModel.Followers {
		follwerIds = append(follwerIds, follower.ID)
	}

	followerStatistcsMaps, err := statistics.GetUsersRelevantCount(l.svcCtx.Redis, l.ctx, follwerIds)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FollowingListLogic] get following statistics list error: %v", err)
	}

	var resp user.FollowerListResp
	for _, follower := range requestModel.Followers {
		followerStatistcs := followerStatistcsMaps[follower.ID]
		resp.List = append(resp.List, &user.User{
			Id:             uint32(follower.ID),
			Username:       follower.Username,
			Nickname:       follower.Nickname,
			Avatar:         follower.Avatar,
			Gender:         uint32(follower.Gender),
			Role:           uint32(follower.Role),
			FollowerCount:  int32(followerStatistcs.FollowerCount),
			FollowingCount: int32(followerStatistcs.FollowingCount),
			LikeCount:      int32(followerStatistcs.LikeCount),
			StarCount:      int32(followerStatistcs.StarCount),
			WorkCount:      int32(followerStatistcs.WorkCount),
			FriendCount:    int32(followerStatistcs.FriendCount),
			Status:         uint32(follower.Status),
		})
	}

	resp.Total = int32(len(resp.List))
	return &resp, nil
}

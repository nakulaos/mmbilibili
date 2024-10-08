package userfollowrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingListLogic {
	return &FollowingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowingListLogic) FollowingList(in *user.FollowingListReq) (*user.FollowingListResp, error) {
	uid := in.ActionId
	if in.Total <= 0 || in.Total > int32(constant.ListDefaultLimit) {
		in.Total = int32(constant.ListDefaultLimit)
	}

	// 获取关注列表
	requestModel, err := l.svcCtx.Dao.User.WithContext(l.ctx).Preload(l.svcCtx.Dao.User.Following.Limit(int(in.Total)).Offset(int(in.PageSize * in.Page))).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FollowingListLogic] get following list error: %v", err)
	}

	var follwingIds []uint
	for _, following := range requestModel.Following {
		follwingIds = append(follwingIds, following.ID)
	}

	// 获取相关数据
	followingStatistcsMaps, err := statistics.GetUsersRelevantCount(l.svcCtx.Redis, l.ctx, follwingIds)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FollowingListLogic] get following statistics list error: %v", err)
	}

	var resp user.FollowingListResp
	for _, following := range requestModel.Following {
		followingStatistcs := followingStatistcsMaps[following.ID]
		resp.List = append(resp.List, &user.User{
			Id:             uint32(following.ID),
			Username:       following.Username,
			Nickname:       following.Nickname,
			Avatar:         following.Avatar,
			Gender:         uint32(following.Gender),
			Role:           uint32(following.Role),
			FollowerCount:  int32(followingStatistcs.FollowerCount),
			FollowingCount: int32(followingStatistcs.FollowingCount),
			LikeCount:      int32(followingStatistcs.LikeCount),
			StarCount:      int32(followingStatistcs.StarCount),
			WorkCount:      int32(followingStatistcs.WorkCount),
			FriendCount:    int32(followingStatistcs.FriendCount),
			Status:         uint32(following.Status),
		})
	}

	resp.Total = int32(len(resp.List))
	return &resp, nil
}

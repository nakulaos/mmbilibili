package userfollowrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *user.FriendListReq) (*user.FriendListResp, error) {
	// 获取用户 ID
	uid := int(in.ActionId)

	// 限制分页参数
	if in.Total <= 0 || in.Total > int32(constant.ListDefaultLimit) {
		in.Total = int32(constant.ListDefaultLimit)
	}

	// 获取自己关注的
	requestModel, err := l.svcCtx.Dao.User.WithContext(l.ctx).Preload(l.svcCtx.Dao.User.Following).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FriendListLogic] get friend list error: %v", err)
	}

	var friendIds []int
	var friendMap = make(map[int]model.User)
	for _, following := range requestModel.Following {
		ret, err := l.svcCtx.Dao.User.Following.WithContext(l.ctx).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).Model(following).Find()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[FriendListLogic] get friend list error: %v", err)
		}
		if len(ret) > 0 {
			friendIds = append(friendIds, int(following.ID))
			friendMap[int(following.ID)] = *following
		}
	}

	sort.Ints(friendIds)

	offset := int(in.PageSize * in.Page)
	if offset >= len(friendMap) {
		return nil, nil
	}

	last := offset + int(in.Total)
	if last > len(friendMap) {
		last = len(friendMap)
	}

	var friendUintIds []uint
	for _, id := range friendIds {
		friendUintIds = append(friendUintIds, uint(id))
	}
	friendStatistcsMaps, err := statistics.GetUsersRelevantCount(l.svcCtx.Redis, l.ctx, friendUintIds)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[FollowingListLogic] get following statistics list error: %v", err)
	}

	var resp user.FriendListResp
	for _, id := range friendIds[offset:last] {
		friend := friendMap[id]
		friendStatistcs := friendStatistcsMaps[uint(id)]
		resp.List = append(resp.List, &user.User{
			Id:       uint32(friend.ID),
			Username: friend.Username,
			Nickname: friend.Nickname,
			Avatar:   friend.Avatar,
			Role:     uint32(friend.Role),

			Gender: uint32(friend.Gender),

			FollowerCount:  int32(friendStatistcs.FollowerCount),
			FollowingCount: int32(friendStatistcs.FollowingCount),
			LikeCount:      int32(friendStatistcs.LikeCount),
			StarCount:      int32(friendStatistcs.StarCount),
			WorkCount:      int32(friendStatistcs.WorkCount),
			FriendCount:    int32(friendStatistcs.FriendCount),
			Status:         uint32(friend.Status),
		})

	}

	resp.Total = int32(len(resp.List))
	return &resp, nil
}

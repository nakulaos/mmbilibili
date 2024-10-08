package follow

import (
	"backend/rpc/user/user"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingListLogic {
	return &FollowingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingListLogic) FollowingList(req *types.FollowingListReq) (*types.FollowingListResp, error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := user.FollowingListReq{
		ActionId: uint32(uid),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		Total:    int32(req.Total),
	}

	followingListResp, err := l.svcCtx.UserFollowServiceClient.FollowingList(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "[FollowingList] call rpc user.FollowingList : userId:%d", uid)
	}

	var resp types.FollowingListResp
	for _, v := range followingListResp.List {
		resp.List = append(resp.List, types.User{
			Id:             v.Id,
			Username:       v.Username,
			Nickname:       v.Nickname,
			Avatar:         v.Avatar,
			Gender:         v.Gender,
			Role:           v.Role,
			FollowerCount:  int(v.FollowerCount),
			FollowingCount: int(v.FollowingCount),
			LikeCount:      int(v.LikeCount),
			StarCount:      int(v.StarCount),
			SelfStarCount:  int(v.SelfStarCount),
			SelfLikeCount:  int(v.SelfLikeCount),
			LiveCount:      int(v.LiveCount),
			WorkCount:      int(v.WorkCount),
			FriendCount:    int(v.FriendCount),
			Phone:          v.Phone,
			Email:          v.Email,
			Status:         uint(v.Status),
		})
	}

	resp.Total = int(followingListResp.Total)
	return &resp, nil
}

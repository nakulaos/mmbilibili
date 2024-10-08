package follow

import (
	"backend/rpc/user/user"
	"context"
	"encoding/json"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := user.FriendListReq{
		ActionId: uint32(uid),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		Total:    int32(req.Total),
	}

	friendListResp, err := l.svcCtx.UserFollowServiceClient.FriendList(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	resp = &types.FriendListResp{
		List: make([]types.User, 0),
	}
	for _, v := range friendListResp.List {
		resp.List = append(resp.List, types.User{
			Id:             v.Id,
			Username:       v.Username,
			Nickname:       v.Nickname,
			Avatar:         v.Avatar,
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
	return resp, nil
}

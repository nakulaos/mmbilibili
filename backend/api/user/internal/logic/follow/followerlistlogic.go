package follow

import (
	"backend/rpc/user/user"
	"context"
	"encoding/json"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (*types.FollowerListResp, error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := user.FollowerListReq{
		ActionId: uint32(uid),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		Total:    int32(req.Total),
	}

	followerListResp, err := l.svcCtx.UserFollowServiceClient.FollowerList(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	var resp types.FollowerListResp
	for _, v := range followerListResp.List {
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

	resp.Total = int(followerListResp.Total)
	return &resp, nil
}

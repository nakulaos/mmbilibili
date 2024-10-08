package common

import (
	"backend/rpc/user/user"
	"context"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.LoginResp, error) {
	in := user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	}
	registerResp, err := l.svcCtx.UserCommonServiceClient.Register(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	resp.AccessToken = registerResp.AccessToken
	userInfo := types.User{
		Id:             registerResp.UserInfo.Id,
		Username:       registerResp.UserInfo.Username,
		Nickname:       registerResp.UserInfo.Nickname,
		Avatar:         registerResp.UserInfo.Avatar,
		Gender:         uint32(registerResp.UserInfo.Gender),
		Role:           uint32(1),
		FollowerCount:  int(registerResp.UserInfo.FollowerCount),
		FollowingCount: int(registerResp.UserInfo.FollowingCount),
		LikeCount:      int(registerResp.UserInfo.LikeCount),
		StarCount:      int(registerResp.UserInfo.StarCount),
		SelfStarCount:  int(registerResp.UserInfo.SelfStarCount),
		SelfLikeCount:  int(registerResp.UserInfo.SelfLikeCount),
		LiveCount:      int(registerResp.UserInfo.LiveCount),
		WorkCount:      int(registerResp.UserInfo.WorkCount),
		FriendCount:    int(registerResp.UserInfo.FriendCount),
		Phone:          registerResp.UserInfo.Phone,
		Email:          registerResp.UserInfo.Email,
	}
	resp.UserInfo = userInfo
	resp.UserID = registerResp.UserId
	return &resp, nil
}

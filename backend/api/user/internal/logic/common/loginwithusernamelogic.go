package common

import (
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithUsernameLogic {
	return &LoginWithUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithUsernameLogic) LoginWithUsername(req *types.LoginWithUsernameReq) (*types.LoginResp, error) {
	in := user.LoginWithUsernameReq{
		Username: req.Username,
		Password: req.Password,
	}

	loginWithUsernameResp, err := l.svcCtx.UserCommonServiceClient.LoginWithUsername(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "[LoginWithUsername] call rpc user.LoginWithUsername : username:%s", req.Username)
	}

	var resp types.LoginResp
	resp.AccessToken = loginWithUsernameResp.AccessToken
	userInfo := types.User{
		Id:             loginWithUsernameResp.UserInfo.Id,
		Username:       loginWithUsernameResp.UserInfo.Username,
		Nickname:       loginWithUsernameResp.UserInfo.Nickname,
		Avatar:         loginWithUsernameResp.UserInfo.Avatar,
		Gender:         uint32(loginWithUsernameResp.UserInfo.Gender),
		Role:           uint32(1),
		FollowerCount:  int(loginWithUsernameResp.UserInfo.FollowerCount),
		FollowingCount: int(loginWithUsernameResp.UserInfo.FollowingCount),
		LikeCount:      int(loginWithUsernameResp.UserInfo.LikeCount),
		StarCount:      int(loginWithUsernameResp.UserInfo.StarCount),
		SelfStarCount:  int(loginWithUsernameResp.UserInfo.SelfStarCount),
		SelfLikeCount:  int(loginWithUsernameResp.UserInfo.SelfLikeCount),
		LiveCount:      int(loginWithUsernameResp.UserInfo.LiveCount),
		WorkCount:      int(loginWithUsernameResp.UserInfo.WorkCount),
		FriendCount:    int(loginWithUsernameResp.UserInfo.FriendCount),
		Phone:          loginWithUsernameResp.UserInfo.Phone,
		Email:          loginWithUsernameResp.UserInfo.Email,
	}
	resp.UserInfo = userInfo
	resp.UserID = loginWithUsernameResp.UserId

	return &resp, nil
}

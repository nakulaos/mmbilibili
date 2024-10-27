package service

import (
	"backend/app/http/biz/global"
	user "backend/app/http/hertz_gen/user"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginWithUsernameService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginWithUsernameService(Context context.Context, RequestContext *app.RequestContext) *LoginWithUsernameService {
	return &LoginWithUsernameService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginWithUsernameService) Run(req *user.LoginWithUsernameReq) (resp *user.LoginResp, err error) {
	loginReq := userRpc.LoginWithUsernameReq{
		Username: req.Username,
		Password: req.Password,
	}

	loginResp, err := global.UserRpcClient.LoginWithUsername(h.Context, &loginReq)

	if err != nil {
		return nil, err
	}

	userInfo := user.User{
		Id:             loginResp.UserInfo.Id,
		Username:       loginResp.UserInfo.Username,
		Nickname:       loginResp.UserInfo.Nickname,
		Avatar:         loginResp.UserInfo.Avatar,
		Gender:         int64(loginResp.UserInfo.Gender),
		Role:           int64(loginResp.UserInfo.Role),
		FollowerCount:  loginResp.UserInfo.FollowerCount,
		FollowingCount: loginResp.UserInfo.FollowingCount,
		LikeCount:      loginResp.UserInfo.LikeCount,
		StarCount:      loginResp.UserInfo.StarCount,
		SelfStarCount:  loginResp.UserInfo.SelfStarCount,
		SelfLikeCount:  loginResp.UserInfo.SelfLikeCount,
		LiveCount:      loginResp.UserInfo.LiveCount,
		WorkCount:      loginResp.UserInfo.WorkCount,
		FriendCount:    loginResp.UserInfo.FriendCount,
		Phone:          loginResp.UserInfo.Phone,
		Email:          loginResp.UserInfo.Email,
	}
	resp = &user.LoginResp{
		AccessToken:  loginResp.AccessToken,
		RefreshToken: loginResp.RefreshToken,
		UserInfo:     &userInfo,
		UserID:       userInfo.Id,
	}

	return
}

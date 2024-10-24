package service

import (
	"backend/app/http/user/biz/dal"
	user "backend/app/http/user/hertz_gen/user"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.LoginResp, err error) {
	registerReq := &userRpc.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	}
	registerResp, err := dal.GlobalUserRpcClient.Register(h.Context, registerReq)

	if err != nil {
		hlog.Errorf("RegisterService.GlobalUserRpcClient.Register(req:%v) error: %v", registerReq, err)
		return nil, err
	}

	userInfo := user.User{
		Id:             registerResp.UserInfo.Id,
		Username:       registerResp.UserInfo.Username,
		Nickname:       registerResp.UserInfo.Nickname,
		Avatar:         registerResp.UserInfo.Avatar,
		Gender:         int64(registerResp.UserInfo.Gender),
		Role:           int64(registerResp.UserInfo.Role),
		FollowerCount:  registerResp.UserInfo.FollowerCount,
		FollowingCount: registerResp.UserInfo.FollowingCount,
		LikeCount:      registerResp.UserInfo.LikeCount,
		StarCount:      registerResp.UserInfo.StarCount,
		SelfStarCount:  registerResp.UserInfo.SelfStarCount,
		SelfLikeCount:  registerResp.UserInfo.SelfLikeCount,
		LiveCount:      registerResp.UserInfo.LiveCount,
		WorkCount:      registerResp.UserInfo.WorkCount,
		FriendCount:    registerResp.UserInfo.FriendCount,
		Phone:          registerResp.UserInfo.Phone,
		Email:          registerResp.UserInfo.Email,
	}
	resp = &user.LoginResp{
		AccessToken:  registerResp.AccessToken,
		RefreshToken: registerResp.RefreshToken,
		UserInfo:     &userInfo,
		UserID:       userInfo.Id,
	}

	return resp, nil
}

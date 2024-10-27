package service

import (
	"backend/app/http/biz/global"
	user "backend/app/http/hertz_gen/user"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	logoutReq := userRpc.LogoutReq{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}

	_, err = global.UserRpcClient.Logout(h.Context, &logoutReq)
	if err != nil {
		hlog.Errorf("LogoutService.GlobalUserRpcClient.Logout(req:%v) error: %v", logoutReq, err)
		return nil, err
	}

	return
}

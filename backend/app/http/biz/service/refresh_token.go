package service

import (
	"backend/app/common/ecode"
	"backend/app/http/biz/global"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RefreshTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRefreshTokenService(Context context.Context, RequestContext *app.RequestContext) *RefreshTokenService {
	return &RefreshTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *RefreshTokenService) Run(req *user.RefreshTokenReq) (resp *user.RefreshTokenResp, err error) {
	refreshTokenAny, exist := h.RequestContext.Get("refresh_token")
	if !exist {
		hlog.Errorf("RefreshTokenService.Get refresh_token error")
		return nil, ecode.ServerError
	}

	refreshToken, ok := refreshTokenAny.(string)
	if !ok {
		hlog.Errorf("RefreshTokenService.Get refresh_token error")
		return nil, ecode.ServerError
	}

	uid := tools.GetUserID(h.RequestContext)
	refreshTokenResp, err := global.UserRpcClient.RefreshToken(h.Context, &userRpc.RefreshTokenReq{
		RefreshToken: refreshToken,
		UserId:       uid,
	})

	if err != nil {
		hlog.Errorf("RefreshTokenService.RefreshToken error:%v", err)
		return nil, ecode.ServerError
	}

	resp = &user.RefreshTokenResp{
		AccessToken:  refreshTokenResp.AccessToken,
		RefreshToken: refreshTokenResp.RefreshToken,
	}

	return resp, nil
}

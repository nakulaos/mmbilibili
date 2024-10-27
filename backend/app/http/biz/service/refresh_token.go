package service

import (
	"context"

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
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

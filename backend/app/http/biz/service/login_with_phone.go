package service

import (
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginWithPhoneService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginWithPhoneService(Context context.Context, RequestContext *app.RequestContext) *LoginWithPhoneService {
	return &LoginWithPhoneService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginWithPhoneService) Run(req *user.LoginWithPhoneReq) (resp *user.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

package service

import (
	"context"

	user "backend/app/http/user/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginWithEmailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginWithEmailService(Context context.Context, RequestContext *app.RequestContext) *LoginWithEmailService {
	return &LoginWithEmailService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginWithEmailService) Run(req *user.LoginWithEmailReq) (resp *user.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

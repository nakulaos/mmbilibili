package service

import (
	"context"

	user "backend/app/http/user/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowerListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowerListService(Context context.Context, RequestContext *app.RequestContext) *FollowerListService {
	return &FollowerListService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowerListService) Run(req *user.FollowerListReq) (resp *user.FollowerListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

package service

import (
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowUserService(Context context.Context, RequestContext *app.RequestContext) *FollowUserService {
	return &FollowUserService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowUserService) Run(req *user.FollowUserReq) (resp *user.FollowUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

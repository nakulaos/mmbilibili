package service

import (
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowingListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowingListService(Context context.Context, RequestContext *app.RequestContext) *FollowingListService {
	return &FollowingListService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowingListService) Run(req *user.FollowingListReq) (resp *user.FollowingListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

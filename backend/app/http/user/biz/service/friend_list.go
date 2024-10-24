package service

import (
	"context"

	user "backend/app/http/user/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FriendListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFriendListService(Context context.Context, RequestContext *app.RequestContext) *FriendListService {
	return &FriendListService{RequestContext: RequestContext, Context: Context}
}

func (h *FriendListService) Run(req *user.FriendListReq) (resp *user.FriendListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

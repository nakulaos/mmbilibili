package service

import (
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserUploadFileService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserUploadFileService(Context context.Context, RequestContext *app.RequestContext) *UserUploadFileService {
	return &UserUploadFileService{RequestContext: RequestContext, Context: Context}
}

func (h *UserUploadFileService) Run(req *user.UserUploadFileReq) (resp *user.UserUploadFileResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}

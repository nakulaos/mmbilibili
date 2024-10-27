package service

import (
	"backend/app/common/ecode"
	"backend/app/http/biz/global"
	fileRpc "backend/app/rpc/file/kitex_gen/file"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	file "backend/app/http/hertz_gen/file"
	"github.com/cloudwego/hertz/pkg/app"
)

type CompleteMultipartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCompleteMultipartService(Context context.Context, RequestContext *app.RequestContext) *CompleteMultipartService {
	return &CompleteMultipartService{RequestContext: RequestContext, Context: Context}
}

func (h *CompleteMultipartService) Run(req *file.CompleteMultipartReq) (resp *file.CompleteMultipartResp, err error) {
	uid := tools.GetUserID(h.RequestContext)
	_, err = global.FileRpcClient.CompleteMultipart(h.Context, &fileRpc.CompleteMultipartReq{
		FileHash: req.FileHash,
		UserID:   uid,
	})

	if err != nil {
		hlog.Errorf("CompleteMultipartService.GlobalFileRpcClient.CompleteMultipart(req:%v) error: %v", req, err)
		return nil, ecode.ServerError
	}
	return &file.CompleteMultipartResp{}, nil
}

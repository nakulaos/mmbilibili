package service

import (
	"backend/app/http/biz/global"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	file "backend/app/http/hertz_gen/file"
	fileRpc "backend/app/rpc/file/kitex_gen/file"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetSuccessChunksService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetSuccessChunksService(Context context.Context, RequestContext *app.RequestContext) *GetSuccessChunksService {
	return &GetSuccessChunksService{RequestContext: RequestContext, Context: Context}
}

func (h *GetSuccessChunksService) Run(req *file.GetSuccessChunksReq) (resp *file.GetSuccessChunksResp, err error) {
	uid := tools.GetUserID(h.RequestContext)
	getSuccessChunksResp, err := global.FileRpcClient.GetSuccessChunks(h.Context, &fileRpc.GetSuccessChunksReq{
		FileHash: req.FileHash,
		UserID:   uid,
	})
	if err != nil {
		hlog.Errorf("GetSuccessChunksService.GlobalFileRpcClient.GetSuccessChunks(req:%v) error: %v", req, err)
		return nil, err
	}
	resp = &file.GetSuccessChunksResp{
		Chunks:   getSuccessChunksResp.Chunks,
		IsUpload: getSuccessChunksResp.IsUpload,
		IsRecord: getSuccessChunksResp.IsRecord,
	}

	return resp, nil
}

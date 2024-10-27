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

type GetMultiUploadUriService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetMultiUploadUriService(Context context.Context, RequestContext *app.RequestContext) *GetMultiUploadUriService {
	return &GetMultiUploadUriService{RequestContext: RequestContext, Context: Context}
}

func (h *GetMultiUploadUriService) Run(req *file.GetMultiUploadUriReq) (resp *file.GetMultiUploadUriResp, err error) {
	uid := tools.GetUserID(h.RequestContext)
	getMultiUploadUriResp, err := global.FileRpcClient.GetMultiUploadUri(h.Context, &fileRpc.GetMultiUploadUriReq{
		FileHash:  req.FileHash,
		UserID:    uid,
		ChunkID:   req.ChunkID,
		ChunkSize: req.ChunkSize,
	})

	if err != nil {
		hlog.Errorf("GetMultiUploadUriService.GlobalFileRpcClient.GetMultiUploadUri(req:%v) error: %v", req, err)
		return nil, ecode.ServerError
	}

	resp = &file.GetMultiUploadUriResp{
		Uri: getMultiUploadUriResp.Uri,
	}

	return resp, nil
}

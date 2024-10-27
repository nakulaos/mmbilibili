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

type NewMultiUploadService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewNewMultiUploadService(Context context.Context, RequestContext *app.RequestContext) *NewMultiUploadService {
	return &NewMultiUploadService{RequestContext: RequestContext, Context: Context}
}

func (h *NewMultiUploadService) Run(req *file.NewMultiUploadReq) (resp *file.NewMultiUploadResp, err error) {
	uid := tools.GetUserID(h.RequestContext)
	_, err = global.FileRpcClient.NewMultiUpload(h.Context, &fileRpc.NewMultiUploadReq{
		FileHash:         req.FileHash,
		UserID:           uid,
		FileName:         req.FileName,
		FileSize:         req.FileSize,
		ChunkTotalNumber: req.ChunkTotalNumber,
		FileType:         req.FileType,
	})

	if err != nil {
		hlog.Errorf("NewMultiUploadService.GlobalFileRpcClient.NewMultiUpload(req:%v) error: %v", req, err)
		return nil, err
	}

	return &file.NewMultiUploadResp{}, nil
}

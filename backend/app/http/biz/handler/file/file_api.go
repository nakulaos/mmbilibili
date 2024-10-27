package file

import (
	"context"

	"backend/app/http/biz/service"
	"backend/app/http/biz/utils"
	file "backend/app/http/hertz_gen/file"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// NewMultiUpload .
// @router /v1/auth/file/multi_upload [POST]
func NewMultiUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req file.NewMultiUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &file.NewMultiUploadResp{}
	resp, err = service.NewNewMultiUploadService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetMultiUploadUri .
// @router /v1/auth/file/multi_upload_uri [POST]
func GetMultiUploadUri(ctx context.Context, c *app.RequestContext) {
	var err error
	var req file.GetMultiUploadUriReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetMultiUploadUriService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CompleteMultipart .
// @router /v1/auth/file/complete_multipart [POST]
func CompleteMultipart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req file.CompleteMultipartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCompleteMultipartService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetSuccessChunks .
// @router /v1/auth/file/success_chunks [POST]
func GetSuccessChunks(ctx context.Context, c *app.RequestContext) {
	var err error
	var req file.GetSuccessChunksReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetSuccessChunksService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

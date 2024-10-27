package main

import (
	"backend/app/rpc/file/biz/service"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct{}

// NewMultiUpload implements the FileServiceImpl interface.
func (s *FileServiceImpl) NewMultiUpload(ctx context.Context, req *file.NewMultiUploadReq) (resp *file.NewMultiUploadResp, err error) {
	resp, err = service.NewNewMultiUploadService(ctx).Run(req)

	return resp, err
}

// GetMultiUploadUri implements the FileServiceImpl interface.
func (s *FileServiceImpl) GetMultiUploadUri(ctx context.Context, req *file.GetMultiUploadUriReq) (resp *file.GetMultiUploadUriResp, err error) {
	resp, err = service.NewGetMultiUploadUriService(ctx).Run(req)

	return resp, err
}

// CompleteMultipart implements the FileServiceImpl interface.
func (s *FileServiceImpl) CompleteMultipart(ctx context.Context, req *file.CompleteMultipartReq) (resp *file.CompleteMultipartResp, err error) {
	resp, err = service.NewCompleteMultipartService(ctx).Run(req)

	return resp, err
}

// GetSuccessChunks implements the FileServiceImpl interface.
func (s *FileServiceImpl) GetSuccessChunks(ctx context.Context, req *file.GetSuccessChunksReq) (resp *file.GetSuccessChunksResp, err error) {
	resp, err = service.NewGetSuccessChunksService(ctx).Run(req)

	return resp, err
}

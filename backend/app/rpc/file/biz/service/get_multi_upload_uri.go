package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/file/biz/global"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type GetMultiUploadUriService struct {
	ctx context.Context
} // NewGetMultiUploadUriService new GetMultiUploadUriService
func NewGetMultiUploadUriService(ctx context.Context) *GetMultiUploadUriService {
	return &GetMultiUploadUriService{ctx: ctx}
}

// Run create note info
func (s *GetMultiUploadUriService) Run(req *file.GetMultiUploadUriReq) (resp *file.GetMultiUploadUriResp, err error) {
	/*
		1. 判定分块文件大小是否超过限制
		2. 生成uri
		3. 返回uri
	*/
	var (
		uri  string
		c    = global.Config
		hash = req.FileHash
		uid  = req.UserID
	)

	fileChunk, err := global.FileDal.GetFileChunkByFileHashANDUserID(s.ctx, hash, uid)
	if err != nil {
		return nil, ecode.ServerError
	}

	var (
		chunkFileSize = req.ChunkSize
	)

	if chunkFileSize >= c.App.ChunkMaxSize*1024*1024 {
		return nil, ecode.FileChunkSizeIllegalError
	}
	buckname := c.MinIO.BucketName
	expire := time.Duration(c.App.PresignedUploadPartUrlExpireTime) * time.Minute
	objectName := fileChunk.ObjectName
	uploadId := fileChunk.UploadID

	if uri, err = global.MinioExtClient.GenUploadPartSignedUrl(
		uploadId, buckname, objectName, int(req.ChunkID), chunkFileSize,
		expire, c.MinIO.Location); err != nil {
		klog.Errorf("getMultiUploadUriService.minioExtClient.genUploadPartSignedUrl(%s,%s) error: %v", buckname, objectName, err)
		return nil, ecode.ServerError
	}

	return &file.GetMultiUploadUriResp{Uri: uri}, nil

}

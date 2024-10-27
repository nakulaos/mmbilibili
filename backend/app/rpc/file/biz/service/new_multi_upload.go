package service

import (
	"backend/app/common/constant"
	"backend/app/common/ecode"
	"backend/app/rpc/file/biz/global"
	"backend/app/rpc/file/biz/model"
	file "backend/app/rpc/file/kitex_gen/file"
	"backend/library/metric"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v6"
	"path/filepath"
)

type NewMultiUploadService struct {
	ctx context.Context
} // NewNewMultiUploadService new NewMultiUploadService
func NewNewMultiUploadService(ctx context.Context) *NewMultiUploadService {
	return &NewMultiUploadService{ctx: ctx}
}

// Run create note info
func (s *NewMultiUploadService) Run(req *file.NewMultiUploadReq) (resp *file.NewMultiUploadResp, err error) {
	/*
		1. 判定文件大小是否超过限制
		2. 生成uuid
		3. 生成objectName
		4. 生成uploadID
		5. 保存文件信息到数据库
		6. 返回uploadID和uuid
	*/

	var (
		uploadID         string
		chunkTotalNumber = req.ChunkTotalNumber
		fileSize         = req.FileSize
		c                = global.Config
		uid              = req.UserID
		hash             = req.FileHash
	)

	if fileSize >= constant.Mb*c.App.FileMaxSize {
		return nil, ecode.FileMaxSizeError
	}

	buckname := c.MinIO.BucketName
	basePath := c.MinIO.BasePath
	ext := filepath.Ext(req.FileName)
	objectName := ObjectName(basePath, fmt.Sprintf("%d", uid), hash, ext)

	// 生成uploadID
	uploadID, err = global.MinioCoreClient.NewMultipartUpload(buckname, objectName, minio.PutObjectOptions{})
	if err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromMinIOUploadID)
		hlog.Errorf("newMultiUploadService.minioCoreClient.newMultipartUpload(%s,%s) error: %v", buckname, objectName, err)
		return nil, ecode.ServerError
	}

	if err = global.FileDal.CreateFileChunk(s.ctx, &model.FileChunk{
		UploadID:    uploadID,
		UserID:      uid,
		TotalChunks: int(chunkTotalNumber),
		FileSize:    fileSize,
		FileName:    req.FileName,
		FileHash:    req.FileHash,
		ObjectName:  objectName,
		FileType:    int16(req.FileType),
	}); err != nil {
		return nil, ecode.ServerError
	}

	resp = &file.NewMultiUploadResp{}

	return resp, nil
}

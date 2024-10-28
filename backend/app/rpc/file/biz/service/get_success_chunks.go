package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/file/biz/global"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"strconv"
)

type GetSuccessChunksService struct {
	ctx context.Context
} // NewGetSuccessChunksService new GetSuccessChunksService
func NewGetSuccessChunksService(ctx context.Context) *GetSuccessChunksService {
	return &GetSuccessChunksService{ctx: ctx}
}

// Run create note info
func (s *GetSuccessChunksService) Run(req *file.GetSuccessChunksReq) (resp *file.GetSuccessChunksResp, err error) {
	var (
		uploadID, chunks string
		uploaded         bool
		fileHash         = req.FileHash
		uid              = req.UserID
		c                = global.Config
	)

	fileChunk, err := global.FileDal.GetFileChunkByFileHashANDUserID(s.ctx, fileHash, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &file.GetSuccessChunksResp{
				IsRecord: false,
			}, nil
		}
		return nil, ecode.ServerError
	}

	uploaded = fileChunk.IsUploaded
	uploadID = fileChunk.UploadID
	bucketName := c.MinIO.BucketName
	objectName := fileChunk.ObjectName

	if fileChunk.IsUploaded {
		resp = &file.GetSuccessChunksResp{
			IsUpload: true,
			IsRecord: true,
		}
		return resp, nil
	}

	uploaded, err = isObjectExist(bucketName, objectName) // 判断文件是否已经上传
	if err != nil {
		klog.Errorf("getSuccessChunksService.isObjectExist(%s,%s) error: %v", bucketName, objectName, err)
		return nil, ecode.ServerError
	}

	if uploaded {
		return &file.GetSuccessChunksResp{
			IsUpload: true,
			IsRecord: true,
		}, nil
	}

	partInfos, err := global.MinioExtClient.ListObjectParts(bucketName, objectName, uploadID)
	if err != nil {
		klog.Errorf("getSuccessChunksService.minioExtClient.listObjectParts(%s,%s,%s) error: %v", bucketName, objectName, uploadID, err)
		return nil, ecode.ServerError
	}

	for _, partInfo := range partInfos {
		chunks += strconv.Itoa(partInfo.PartNumber) + ","
	}

	resp = &file.GetSuccessChunksResp{
		IsUpload: false,
		IsRecord: true,
		Chunks:   chunks,
	}

	return
}

func isObjectExist(bucketName string, objectName string) (bool, error) {
	var (
		isExist bool
	)
	doneCh := make(chan struct{})
	defer close(doneCh)

	objectCh := global.MinioClient.ListObjects(bucketName, objectName, false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			klog.Errorf("isObjectExist.minioClient.listObjects(%s,%s) error: %v", bucketName, objectName, object.Err)
			return isExist, object.Err
		}
		isExist = true
		break
	}

	return isExist, nil
}

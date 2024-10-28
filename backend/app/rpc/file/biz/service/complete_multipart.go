package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/file/biz/global"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"encoding/xml"
	"github.com/cloudwego/kitex/pkg/klog"
	miniov6 "github.com/minio/minio-go/v6"
	"sort"
)

type CompleteMultipartService struct {
	ctx context.Context
} // NewCompleteMultipartService new CompleteMultipartService
func NewCompleteMultipartService(ctx context.Context) *CompleteMultipartService {
	return &CompleteMultipartService{ctx: ctx}
}

// Run create note info
func (s *CompleteMultipartService) Run(req *file.CompleteMultipartReq) (resp *file.CompleteMultipartResp, err error) {
	var (
		userID   = req.UserID
		filehash = req.FileHash
		c        = global.Config
		uploadID string
	)

	fileChunk, err := global.FileDal.GetFileChunkByFileHashANDUserID(s.ctx, filehash, userID)
	if err != nil {
		return nil, ecode.ServerError
	}
	uploadID = fileChunk.UploadID
	// complete multipart upload
	bucketName := c.MinIO.BucketName
	objectName := fileChunk.ObjectName
	partInfos, err := global.MinioExtClient.ListObjectParts(bucketName, objectName, uploadID)
	if err != nil {
		klog.Errorf("completeMultipartService.minioExtClient.listObjectParts(%s,%s,%s) error: %v", bucketName, objectName, uploadID, err)
		return nil, ecode.ServerError
	}
	var complMultipartUpload completeMultipartUpload
	for _, partInfo := range partInfos {
		complMultipartUpload.Parts = append(complMultipartUpload.Parts, miniov6.CompletePart{
			PartNumber: partInfo.PartNumber,
			ETag:       partInfo.ETag,
		})
	}
	sort.Sort(completedParts(complMultipartUpload.Parts))
	_, err = global.MinioCoreClient.CompleteMultipartUpload(bucketName, objectName, uploadID, complMultipartUpload.Parts)
	if err != nil {
		klog.Errorf("completeMultipartService.minioCoreClient.completeMultipartUpload(%s,%s,%s) error: %v", bucketName, objectName, uploadID, err)
		return nil, err
	}

	// db
	fileChunk.IsUploaded = true

	if err = global.FileDal.SaveFileChunk(s.ctx, fileChunk); err != nil {
		return nil, ecode.ServerError
	}

	return &file.CompleteMultipartResp{}, nil
}

type completeMultipartUpload struct {
	XMLName xml.Name               `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CompleteMultipartUpload" json:"-"`
	Parts   []miniov6.CompletePart `xml:"Part"`
}

type completedParts []miniov6.CompletePart

func (a completedParts) Len() int           { return len(a) }
func (a completedParts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a completedParts) Less(i, j int) bool { return a[i].PartNumber < a[j].PartNumber }

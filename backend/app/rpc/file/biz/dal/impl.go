package dal

import (
	"backend/app/common/constant"
	"backend/app/rpc/file/biz/model"
	"backend/library/metric"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

var _ FileDal = (*FileDalImpl)(nil)

type FileDalImpl struct {
	db *gorm.DB
}

func (f FileDalImpl) SaveFileChunk(ctx context.Context, fileChunk *model.FileChunk) error {
	if err := f.db.Save(fileChunk).Error; err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBFileChunk)
		klog.Errorf("db.save(%v) error: %v", fileChunk, err)
	}
	return nil
}

func (f FileDalImpl) GetFileChunkByFileHashANDUserID(ctx context.Context, fileHash string, userID int64) (*model.FileChunk, error) {
	fileChunk := &model.FileChunk{}
	if err := f.db.Where("file_hash = ? AND user_id = ?", fileHash, userID).First(fileChunk).Error; err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBFileChunk)
		klog.Errorf("db.where(%s,%d).first error: %v", fileHash, userID, err)
		return nil, err
	}
	return fileChunk, nil
}

func (f FileDalImpl) CreateFileChunk(ctx context.Context, fileChunk *model.FileChunk) error {
	if err := f.db.Create(fileChunk).Error; err != nil {
		metric.IncrGauge(metric.LibClient, constant.PromDBFileChunk)
		klog.Errorf("db.create(%v) error: %v", fileChunk, err)
		return err
	}
	return nil
}

func NewFileDalImpl(db *gorm.DB) *FileDalImpl {
	return &FileDalImpl{db: db}
}

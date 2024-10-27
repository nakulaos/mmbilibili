package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	FileTypeImage = iota + 1
	FileTypeVideo
	FileTypeAudio
	FileTypeDocument
	FileTypeOther
)

const (
	FileNotUploaded int = iota
	FileUploaded
)

type FileChunk struct {
	ID          int64          `gorm:"primaryKey;autoIncrement"` // 自定义ID字段
	CreatedAt   time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:删除时间"`         // 软删除
	FileHash    string         `gorm:"UNIQUE"`                     // user_id + file_hash 文件hash,去确定文件是否已经上传
	IsUploaded  int            `gorm:"DEFAULT 0"`                  // not uploaded: 0, uploaded: 1
	UploadID    string         `gorm:"UNIQUE"`                     //minio upload id
	TotalChunks int            `gorm:"not null"`                   // 总块数
	FileSize    int64          `gorm:"not null"`                   // 文件大小
	FileName    string         `gorm:"type:varchar(255);not null"` // 文件名
	ObjectName  string         `gorm:"type:varchar(255);not null"` // minio object name
	UserID      int64          `gorm:"not null" json:"user_id"`    // 用户ID,那一个用户上传的
	FileType    int16          `gorm:"not null" json:"file_type"`  // 文件类型
}

func (FileChunk) TableName() string {
	return "file_chunks"
}

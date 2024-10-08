package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null;comment:标签名;default:'';unique_index"` // 标签名
	Description string `gorm:"type:text"`                                                      // 增加描述
}

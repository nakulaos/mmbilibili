package model

import (
	"gorm.io/gorm"
	"time"
)

type UserRelationship struct {
	ID               int64          `gorm:"primaryKey;autoIncrement"` // 自定义ID字段
	CreatedAt        time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt        time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt        gorm.DeletedAt `gorm:"index;comment:删除时间"`                          // 软删除
	UserID           int64          `gorm:"not null" json:"user_id"`                     // 用户ID
	RelatedUserID    int64          `gorm:"not null" json:"related_user_id"`             // 相关用户ID
	RelationshipAttr int64          `gorm:"not null;default:0" json:"relationship_attr"` // 关系属性
}

func (u *UserRelationship) TableName() string {
	return "user_relationships"
}

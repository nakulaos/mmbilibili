package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	AttrNormalState = iota + 1
	AttrBanLoginState
)

type User struct {
	ID          int64          `gorm:"primaryKey;autoIncrement"` // 自定义ID字段
	CreatedAt   time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:删除时间"`                                           // 软删除
	Role        uint8          `gorm:"type:tinyint unsigned;default:1;not null"`                     // 角色，默认值为1
	Gender      uint8          `gorm:"type:tinyint unsigned;default:1;not null"`                     // 性别，默认值为1
	Status      uint8          `gorm:"type:tinyint unsigned;default:1;not null;comment:'1:正常 2:封禁'"` // 1:正常 2:封禁
	Username    string         `gorm:"type:varchar(48);not null;unique;comment:账号名"`                 // 唯一约束
	Nickname    string         `gorm:"type:varchar(48);not null"`
	Description string         `gorm:"type:varchar(256);default:null;comment:简介"`
	Salt        string         `gorm:"type:varchar(36);not null"`
	Phone       string         `gorm:"type:varchar(48);unique;default:null"` // 可以为空，且唯一
	Email       string         `gorm:"type:varchar(48);unique;default:null"` // 可以为空，且唯一
	Password    string         `gorm:"type:varchar(256);not null"`
	Avatar      string         `gorm:"type:varchar(255);not null"`
}

// TableName 为 User 模型指定数据库表名
func (u *User) TableName() string {
	return "users"
}

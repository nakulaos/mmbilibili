/**
 ******************************************************************************
 * @file           : live.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

// Live struct
// 这个就是一个直播记录表
type Live struct {
	UID          uint       `gorm:"not null;comment:用户ID"`
	Title        string     `gorm:"type:varchar(255);not null;comment:直播标题"`
	Description  string     `gorm:"type:varchar(255);comment:直播描述"`
	StartTime    time.Time  `gorm:"not null;comment:开始时间"`
	EndTime      time.Time  `gorm:"comment:结束时间"`
	Status       uint8      `gorm:"not null;default:0;comment:直播状态"`
	Danmus       []Danmu    `gorm:"polymorphic:Owner;"`
	Categories   []Category `gorm:"many2many:live_category;"`
	Tags         []Tag      `gorm:"many2many:live_tag;"`
	PlayURL      string     `gorm:"type:varchar(255);not null;comment:直播地址"`
	CoverURL     string     `gorm:"type:varchar(255);not null;comment:封面地址"`
	StreamID     uint       `gorm:"not null;default:0;comment:流ID"`
	IsOver       int        `gorm:"not null;default:0;comment:是否结束"`
	RoomID       uint       `gorm:"not null;default:0;comment:房间ID"`
	PushToken    string     `gorm:"type:varchar(255);comment:推流token"`
	PlayToken    string     `gorm:"type:varchar(255);comment:播放token"`
	Partition    string     `gorm:"type:varchar(255);comment:分区"`
	FavoriteUser []User     `gorm:"many2many:user_favorite_lives;"`
	//ViewCount    uint       `gorm:"not null;default:0;comment:观看人数"`
	//LikeCount    uint       `gorm:"not null;default:0;comment:点赞数"`
	//CommentCount uint       `gorm:"not null;default:0;comment:评论数"`
	gorm.Model
}

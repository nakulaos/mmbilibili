/**
 ******************************************************************************
 * @file           : video.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

// Video struct
type Video struct {
	gorm.Model
	AuthorID   uint       `gorm:"not null;comment:上传用户Id"`
	Title      string     `gorm:"type:varchar(255);not null;comment:视频标题"`
	CoverURL   string     `gorm:"type:varchar(255);not null;comment:封面url"`
	PlayURL    string     `gorm:"type:varchar(255);not null;comment:视频播放url"`
	Category   uint       `gorm:"not null;comment:视频分类"`
	Duration   string     `gorm:"type:varchar(255);comment:视频时长"`
	Comments   []Comment  `gorm:"polymorphic:Owner;"`
	Danmu      []Danmu    `gorm:"polymorphic:Owner;"`
	Tags       []Tag      `gorm:"many2many:video_tag;"` // 视频标签
	Categories []Category `gorm:"many2many:video_category;"`
	Status     uint       `gorm:"not null;default:1;comment:视频状态"`
	Copyright  string     `gorm:"type:varchar(255);comment:视频版权"`
	Visibility uint       `gorm:"not null;default:1;comment:视频可见性"`
	//FavoriteCount uint       `gorm:"not null;default:0;comment:点赞数"`
	//StarCount     uint       `gorm:"not null;comment:收藏数"`
	//CommentCount  uint       `gorm:"not null;default:0;comment:评论数目"`
}

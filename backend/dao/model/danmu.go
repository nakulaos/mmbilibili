/**
 ******************************************************************************
 * @file           : danmu.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

// Danmu struct
type Danmu struct {
	gorm.Model
	UID         uint    `gorm:"not null"`
	OwnerID     uint    `gorm:"not null;index:idx_owner,priority:1"` // OwnerID
	OwnerType   string  `gorm:"not null;index:idx_owner,priority:2"` // OwnerType
	Content     string  `gorm:"type:text;not null"`
	SendTime    float64 `gorm:"not null;comment:在视频的哪个点发送的弹幕"`
	Type        int     `gorm:"not null;comment:弹幕类型，1 为视频，2 为直播"`
	LikeCount   int     `gorm:"default:0"`
	UnLikeCount int     `gorm:"default:0"`
}

/**
 ******************************************************************************
 * @file           : comment.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

// Comment struct with polymorphic fields
type Comment struct {
	gorm.Model
	UID         uint      `gorm:"not null"`
	OwnerID     uint      `gorm:"not null"` // Polymorphic ID
	OwnerType   string    `gorm:"not null"` // Polymorphic Type
	Content     string    `gorm:"type:text;not null"`
	ParentID    *uint     `gorm:"index"`               // Self-referencing foreign key for parent comment
	Replies     []Comment `gorm:"foreignKey:ParentID"` // List of child comments
	LikeCount   int       `gorm:"default:0"`
	UnLikeCount int       `gorm:"default:0"`
}

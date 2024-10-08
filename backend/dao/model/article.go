/**
 ******************************************************************************
 * @file           : article.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

// Article struct
type Article struct {
	gorm.Model
	Title      string     `gorm:"type:varchar(255);not null;comment:文章标题"`
	Content    string     `gorm:"type:text;not null;comment:文章内容"`
	AuthorID   uint       `gorm:"not null;comment:作者ID"`
	Categories []Category `gorm:"many2many:article_categories;"`
	Comments   []Comment  `gorm:"polymorphic:Owner;"`
	Tags       []Tag      `gorm:"many2many:article_tags;"`
}

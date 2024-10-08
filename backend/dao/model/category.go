/**
 ******************************************************************************
 * @file           : category.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

// Category 即是种类，也是分区
// 网游，手游 ，
type Category struct {
	Type        string `gorm:"type:varchar(255);not null;index:idx_name_type,unique"`
	Partition   string `gorm:"type:varchar(255);not null;index:idx_name_type,unique"`
	Name        string `gorm:"type:varchar(255);not null;comment:分类名;default:'';index:idx_name_type,unique"` // 分类名
	Description string `gorm:"type:text"`                                                                    // 增加描述
	Tags        []Tag  `gorm:"many2many:category_tags"`
	gorm.Model
}

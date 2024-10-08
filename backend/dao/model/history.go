/**
 ******************************************************************************
 * @file           : history.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UID    uint `gorm:"not null"`
	WorkID uint `gorm:"not null"`
	Type   int  `gorm:"not null;comment:1: video, 2: live 3. article"`
}

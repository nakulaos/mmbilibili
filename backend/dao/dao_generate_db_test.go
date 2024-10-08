/**
 ******************************************************************************
 * @file           : dao_generate_db_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package dao

import (
	"backend/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestDBAutoMigerate(t *testing.T) {
	dsn := "root:asdasd@tcp(127.0.0.1:13306)/cctiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(
		&model.User{},
		&model.Video{},
		&model.Comment{},
		&model.Article{},
		&model.Live{},
		&model.Danmu{},
		&model.Category{},
		&model.History{},
		// Add other models here
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}

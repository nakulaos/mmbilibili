package initializer

import (
	globalConf "backend/library/conf"
	"go.opentelemetry.io/otel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"time"
)

func InitGormDBFromMysql(c globalConf.Mysql) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Minute)

	if err = db.Use(tracing.NewPlugin(tracing.WithTracerProvider(otel.GetTracerProvider()))); err != nil {
		panic(err)
	}

	return db
}

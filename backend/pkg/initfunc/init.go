package initfunc

import (
	"backend/pkg/config"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	redis3 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitGormDB(c config.Mysql) (*gorm.DB, error) {
	mysqlConf := c
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
	// 添加缓存

	// Open the database connection using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logx.WithContext(context.Background()).Error("failed to connect database: %v", err)
		return nil, err
	}
	return db, nil
}

func InitGorseClient(recommenduri string) *client.GorseClient {
	gorseClient := client.NewGorseClient(recommenduri, "")
	return gorseClient
}

func InitCache(c redis3.RedisConf) *cache.Cache {
	redisClient := redis2.NewClient(&redis2.Options{
		Addr: c.Host,
	})

	cache := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1500000, time.Minute),
	})
	return cache
}

//
//func InitGormDBWithCache(sqlConf config.Mysql, redisConf redis2.RedisConf) (*gorm.DB, error) {
//	mysqlConf := sqlConf
//	cacheConf := redisConf
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
//
//	redisClient, err := redis2.NewRedis(cacheConf)
//	if err != nil {
//		logx.WithContext(context.Background()).Error("failed to connect redis: %v", err)
//		return nil, err
//	}
//	cache, _ := cache.NewGorm2Cache(&config2.CacheConfig{
//		CacheLevel:           config2.CacheLevelAll,
//		CacheStorage:         config2.CacheStorageRedis,
//		RedisConfig:          cache.NewRedisConfigWithClient(redisClient),
//		InvalidateWhenUpdate: true,    // when you create/update/delete objects, invalidate cache
//		CacheTTL:             5000000, // 5000 ms
//		CacheMaxItemCnt:      5,       // if length of objects retrieved one single time
//		// exceeds this number, then don't cache
//		//DebugMode: true,
//	})
//
//	// Open the database connection using GORM
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	db.Use(cache)
//
//	if err != nil {
//		logx.WithContext(context.Background()).Error("failed to connect database: %v", err)
//		return nil, err
//	}
//
//	return db, nil
//}

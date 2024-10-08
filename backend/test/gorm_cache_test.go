package test

import (
	"backend/3rdpart/gorm-cache/cache"
	"backend/3rdpart/gorm-cache/config"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	ID    int
	Value int
}

func TestGormCache(t *testing.T) {
	dsn := "root:asdasd@tcp(127.0.0.1:13306)/test?charset=utf8mb4"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	redisClient, err := redis.NewRedis(redis.RedisConf{
		Host: "localhost:36379",
		Type: "node",
	})
	if err != nil {
		t.Fatalf("Failed to create Redis client: %v", err)
	}

	//redisClient := redis.NewClusterClient()

	cache, _ := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageRedis,
		RedisConfig:          cache.NewRedisConfigWithClient(redisClient),
		InvalidateWhenUpdate: true,   // when you create/update/delete objects, invalidate cache
		CacheTTL:             500000, // 5000 ms
		CacheMaxItemCnt:      5,      // if length of objects retrieved one single time
		// exceeds this number, then don't cache
		DebugMode: true,
	})
	// More options in `config.config.go`
	db.Use(cache) // use gorm plugin
	// cache.AttachToDB(db)

	var users []User

	db.Where("value > ?", 123).Find(&users) // search cache not hit, objects cached
	db.Where("value > ?", 123).Find(&users) // search cache hit

	db.Where("id IN (?)", []int{1, 2, 3}).Find(&users) // primary key cache not hit, users cached
	var user User
	user.ID = 1
	user.Value = 123
	db.Save(user)
	db.Where("id IN (?)", []int{1, 3}).Find(&users) // primary key cache hit
	fmt.Println(users)

	user.ID = 2
	user.Value = 456
	db.Save(user)
}

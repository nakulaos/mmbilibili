package svc

import (
	"backend/dao/query"
	"backend/pkg/initfunc"
	"backend/pkg/lang"
	"backend/rpc/user/internal/config"
	"github.com/go-redis/cache/v9"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	Dao         *query.Query
	DB          *gorm.DB
	Cache       *cache.Cache
	Redis       *redis.Redis
	GorseClient *client.GorseClient
	Bundle      *i18n.Bundle
	//KqRegisterUserPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := initfunc.InitGormDB(c.Mysql)
	if err != nil {
		panic(err)
	}

	redis, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:      c,
		DB:          db,
		Dao:         query.Use(db),
		Cache:       initfunc.InitCache(c.RedisConf),
		Redis:       redis,
		GorseClient: initfunc.InitGorseClient(c.App.RecommendUrl),
		Bundle:      lang.GetBundle(),
		//KqRegisterUserPusher: kq.NewPusher(c.KqRegisterPush.Brokers, c.KqRegisterPush.Topic, kq.WithAllowAutoTopicCreation()),
	}
}

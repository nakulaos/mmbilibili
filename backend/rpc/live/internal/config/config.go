package config

import (
	"backend/pkg/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql     config.Mysql
	RedisConf redis.RedisConf
	App       config.App
	//KqAddLivePush config.KqPusher
}

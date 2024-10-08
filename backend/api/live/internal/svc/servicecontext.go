package svc

import (
	"backend/api/live/internal/config"
	"backend/api/live/internal/middleware"
	"backend/rpc/live/live"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                    config.Config
	CorsMiddleware            rest.Middleware
	AuthMiddleware            rest.Middleware
	ApiHitMiddleware          rest.Middleware
	ErrorcodeMiddleware       rest.Middleware
	Redis                     *redis.Redis
	LiveBusinessServiceClient live.LiveBusinessRpcServiceClient
	LiveCallbackServiceClient live.LiveCallbackRpcServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis, err := redis.NewRedis(c.RedisConf)

	if err != nil {
		panic(err)
	}

	liveRpc, err := zrpc.NewClient(c.LiveRpc)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                    c,
		CorsMiddleware:            middleware.NewCorsMiddleware().Handle,
		AuthMiddleware:            middleware.NewAuthMiddleware(redis).Handle,
		ApiHitMiddleware:          middleware.NewApiHitMiddleware().Handle,
		ErrorcodeMiddleware:       middleware.NewErrorcodeMiddleware().Handle,
		Redis:                     redis,
		LiveBusinessServiceClient: live.NewLiveBusinessRpcServiceClient(liveRpc.Conn()),
		LiveCallbackServiceClient: live.NewLiveCallbackRpcServiceClient(liveRpc.Conn()),
	}
}

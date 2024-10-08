package svc

import (
	"backend/api/user/internal/config"
	"backend/api/user/internal/middleware"
	"backend/rpc/user/client/userfollowrpcservice"
	"backend/rpc/user/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                  config.Config
	CorsMiddleware          rest.Middleware
	ApiHitMiddleware        rest.Middleware
	ErrorcodeMiddleware     rest.Middleware
	AuthMiddleware          rest.Middleware
	UserCommonServiceClient user.UserCommonRpcServiceClient
	UserFollowServiceClient user.UserFollowRpcServiceClient
	UserFileServiceClient   user.UserFileRpcServiceClient
	Redis                   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpcConn, err := zrpc.NewClient(c.UserRpc)
	if err != nil {
		panic(err)
	}
	redis, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                  c,
		CorsMiddleware:          middleware.NewCorsMiddleware().Handle,
		ApiHitMiddleware:        middleware.NewApiHitMiddleware().Handle,
		ErrorcodeMiddleware:     middleware.NewErrorcodeMiddleware().Handle,
		AuthMiddleware:          middleware.NewAuthMiddleware(redis).Handle,
		UserCommonServiceClient: user.NewUserCommonRpcServiceClient(userRpcConn.Conn()),
		UserFollowServiceClient: userfollowrpcservice.NewUserFollowRpcService(userRpcConn),
		UserFileServiceClient:   user.NewUserFileRpcServiceClient(userRpcConn.Conn()),
		Redis:                   redis,
	}
}

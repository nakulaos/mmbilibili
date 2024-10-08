package main

import (
	"backend/pkg/interceptor"
	"flag"
	"fmt"

	"backend/rpc/user/internal/config"
	usercommonrpcserviceServer "backend/rpc/user/internal/server/usercommonrpcservice"
	userfilerpcserviceServer "backend/rpc/user/internal/server/userfilerpcservice"
	userfollowrpcserviceServer "backend/rpc/user/internal/server/userfollowrpcservice"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserCommonRpcServiceServer(grpcServer, usercommonrpcserviceServer.NewUserCommonRpcServiceServer(ctx))
		user.RegisterUserFollowRpcServiceServer(grpcServer, userfollowrpcserviceServer.NewUserFollowRpcServiceServer(ctx))
		user.RegisterUserFileRpcServiceServer(grpcServer, userfilerpcserviceServer.NewUserFileRpcServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.ServerLogInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

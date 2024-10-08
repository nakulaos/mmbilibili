package main

import (
	"backend/pkg/interceptor"
	"flag"
	"fmt"

	"backend/rpc/live/internal/config"
	livebusinessrpcserviceServer "backend/rpc/live/internal/server/livebusinessrpcservice"
	livecallbackrpcserviceServer "backend/rpc/live/internal/server/livecallbackrpcservice"
	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/live.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		live.RegisterLiveBusinessRpcServiceServer(grpcServer, livebusinessrpcserviceServer.NewLiveBusinessRpcServiceServer(ctx))
		live.RegisterLiveCallbackRpcServiceServer(grpcServer, livecallbackrpcserviceServer.NewLiveCallbackRpcServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.ServerLogInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

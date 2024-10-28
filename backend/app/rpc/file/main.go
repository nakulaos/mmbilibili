package main

import (
	"backend/app/rpc/file/biz/global"
	"backend/app/rpc/file/kitex_gen/file/fileservice"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	global.MustInitGlobalVal()

	// 服务注册
	r, err := consul.NewConsulRegister(global.Config.Registry.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 链路追踪+metric
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(global.Config.Kitex.Service),
		provider.WithExportEndpoint(global.Config.OpenTelemetry.OpenTelemetryCollectorAddr),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	addr, err := net.ResolveTCPAddr("tcp", global.Config.Kitex.Address)
	if err != nil {
		panic(err)
	}
	var opts []server.Option = make([]server.Option, 0)

	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: global.Config.Kitex.Service,
		}),
		server.WithRegistry(r),
		server.WithSuite(tracing.NewServerSuite()),
	)

	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(global.LogLevel(global.Config.Kitex.Log.LogLevel))
	//asyncWriter := &zapcore.BufferedWriteSyncer{
	//	WS: zapcore.AddSync(&lumberjack.Logger{
	//		Filename:   global.Config.Kitex.Log.LogFileName,
	//		MaxSize:    global.Config.Kitex.Log.LogMaxSize,
	//		MaxBackups: global.Config.Kitex.Log.LogMaxBackups,
	//		MaxAge:     global.Config.Kitex.Log.LogMaxAge,
	//	}),
	//	FlushInterval: time.Minute,
	//}
	//klog.SetOutput(asyncWriter)
	//server.RegisterShutdownHook(func() {
	//	asyncWriter.Sync()
	//})

	svr := fileservice.NewServer(new(FileServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

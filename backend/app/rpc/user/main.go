package main

import (
	"backend/app/rpc/user/biz/dal"
	"backend/app/rpc/user/conf"
	"backend/app/rpc/user/kitex_gen/user/userrpcservice"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net"
	"time"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	c := conf.GetConf()
	r, err := consul.NewConsulRegister(c.Registry.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(c.Kitex.Service),
		provider.WithExportEndpoint(c.OpenTelemetry.OpenTelemetryCollectorAddr),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	addr, err := net.ResolveTCPAddr("tcp", c.Kitex.Address)
	if err != nil {
		panic(err)
	}
	var opts []server.Option = make([]server.Option, 0)

	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: c.Kitex.Service,
		}),
		server.WithRegistry(r),
		server.WithSuite(tracing.NewServerSuite()),
	)

	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.Log.LogFileName,
			MaxSize:    conf.GetConf().Kitex.Log.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.Log.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.Log.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})

	dal.Init(*c)
	svr := userrpcservice.NewServer(new(UserRpcServiceImpl), opts...)

	//go func() {
	//	for {
	//		time.Sleep(3 * time.Second)
	//		for i := 0; i < 5; i++ {
	//			metric.IncrGauge(metric.BusinessInfo, "test.kitex.metric")
	//		}
	//	}
	//}()

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

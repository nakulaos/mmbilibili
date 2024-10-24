// Code generated by hertz generator.

package main

import (
	"backend/app/http/user/biz/dal"
	"backend/app/http/user/biz/router"
	"backend/app/http/user/conf"
	"backend/library/initializer"
	"backend/library/tools"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"log"
)

func main() {

	c := conf.GetConf()

	// 链路追踪
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(c.Hertz.Service),
		provider.WithExportEndpoint(c.OpenTelemetry.OpenTelemetryCollectorAddr),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertztracing.NewServerTracer()

	// 服务注册
	r := initializer.InitHertzRegistry(c.Registry.RegistryAddress)
	localIP, err := tools.GetLocalIPv4()
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s%s", localIP, c.Hertz.Address)
	h := server.New(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: c.Hertz.Service,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      c.Registry.Weight,
			Tags:        nil,
		}),
		server.WithValidateConfig(initializer.InitValidator()),
		tracer)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	initializer.MustInitMiddleware(h)
	h.Use(hertztracing.ServerMiddleware(cfg))
	initializer.MustInitLogger(h, c.Hertz.Log.LogFileName, c.Hertz.Log.LogLevel, c.Hertz.Log.LogMaxSize, c.Hertz.Log.LogMaxBackups, c.Hertz.Log.LogMaxAge)
	dal.Init(*c)
	router.GeneratedRegister(h)

	h.Spin()
}

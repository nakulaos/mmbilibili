package global

import (
	"backend/app/http/conf"
	"backend/app/rpc/file/kitex_gen/file/fileservice"
	"backend/app/rpc/user/kitex_gen/user/userrpcservice"
	"backend/library/initializer"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"time"
)

func MustInitGlobal(c conf.Config) {
	r, err := consul.NewConsulResolver(c.Resolve.ResolveAddress)
	rpcConf := c.UserRpc
	if err != nil {
		klog.Fatal("new consul resolver error: %v", err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(rpcConf.Name),
		provider.WithExportEndpoint(c.OpenTelemetry.OpenTelemetryCollectorAddr),
		provider.WithInsecure(),
	)

	UserRpcClient = userrpcservice.MustNewClient(c.UserRpc.Name,
		client.WithRPCTimeout(time.Duration(c.UserRpc.RPCTimeout)*time.Second),
		client.WithConnectTimeout(time.Duration(c.UserRpc.ConnTimeout)*time.Second),
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcConf.Name}),
	)

	FileRpcClient = fileservice.MustNewClient(c.FileRpc.Name,
		client.WithRPCTimeout(time.Duration(c.FileRpc.RPCTimeout)*time.Second),
		client.WithConnectTimeout(time.Duration(c.FileRpc.ConnTimeout)*time.Second),
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.FileRpc.Name}),
	)

	RedisClient = initializer.InitRedisUniversal(
		c.Redis.Addrs,
		c.Redis.ClientName,
		c.Redis.DialTimeout,
		c.Redis.ReadTimeout,
		c.Redis.WriteTimeout,
		c.Redis.MaxActiveCoons,
		c.Redis.MinIdleCoons)

}
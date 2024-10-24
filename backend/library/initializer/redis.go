package initializer

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"time"
)

func InitRedisUniversal(addr []string, clientName string, dialTimeout, readTimeout, writeTimeout, maxActiveCoons, minIdleCoons int) redis.UniversalClient {
	r := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:          addr,
		ClientName:     clientName,
		DialTimeout:    time.Second * time.Duration(dialTimeout),
		ReadTimeout:    time.Second * time.Duration(readTimeout),
		WriteTimeout:   time.Second * time.Duration(writeTimeout),
		MaxActiveConns: maxActiveCoons, // 最大活跃连接数
		MinIdleConns:   minIdleCoons,   // 最小空闲连接数
	})
	if err := r.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	if err := redisotel.InstrumentTracing(r, redisotel.WithTracerProvider(otel.GetTracerProvider())); err != nil {
		panic(err)
	}

	if err := redisotel.InstrumentMetrics(r, redisotel.WithMeterProvider(otel.GetMeterProvider())); err != nil {
		panic(err)
	}
	return r
}

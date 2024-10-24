package metric

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// MeterType 定义度量的类型
type MeterType uint8

var (
	globalMeter = otel.Meter("global_meter")
	m           sync.Map // 不需要使用 make，直接声明即可
)

// MeterType 的常量定义
const (
	CacheHit MeterType = iota + 1
	CacheMiss
	BusinessInfo
	BusinessError
	LibClient
	RPCClient
	HTTPClient
	HTTPServer
	RPCServer
)

// CacheCallback 定义缓存回调函数类型
type CacheCallback func()

// toMeterType 将 MeterType 转换为字符串
func toMeterType(t MeterType) string {
	switch t {
	case CacheHit:
		return "cache_hit."
	case CacheMiss:
		return "cache_miss."
	case BusinessInfo:
		return "business_info."
	case BusinessError:
		return "business_error."
	case LibClient:
		return "lib_client."
	case RPCClient:
		return "rpc_client."
	case HTTPClient:
		return "http_client."
	case HTTPServer:
		return "http_server."
	case RPCServer:
		return "rpc_server."
	default:
		return "unknown type."
	}
}

// newGauge 创建新的仪表盘并返回对应的回调函数
func newGauge(type_ MeterType, name, description string) CacheCallback {
	var cnt atomic.Int64
	_, err := globalMeter.Int64ObservableGauge(
		toMeterType(type_)+name,
		metric.WithDescription(description),
		metric.WithUnit("1"),
		metric.WithInt64Callback(func(_ context.Context, observer metric.Int64Observer) error {
			observer.Observe(cnt.Load())
			cnt.Store(0)
			return nil
		}),
	)
	if err != nil {
		hlog.Error("failed to create observable gauge: %v", err)
	}

	cacheCallback := CacheCallback(
		func() {
			cnt.Add(1)
		})

	m.Store(toMeterType(type_)+name, cacheCallback)

	return cacheCallback
}

// IncrGauge 增加指定度量的计数
func IncrGauge(type_ MeterType, name string) {
	key := toMeterType(type_) + name
	if v, ok := m.Load(key); ok {
		if callback, ok := v.(CacheCallback); ok {
			callback()
		} else {
			hlog.Errorf("callback for %s is not of type CacheCallback", key)
		}
	} else {
		hlog.Infof("not found %s, creating new gauge", key)
		newGauge(type_, name, "")
		hlog.Infof("created %s", key)
		IncrGauge(type_, name)
	}
}

func CleanMetrics() {
	m.Range(func(key, value interface{}) bool {
		m.Delete(key) // 删除所有条目
		return true   // 继续迭代
	})
	hlog.Infof("cleaned all metrics")
}

func RegisterGauge(type_ MeterType, name, description string) {
	newGauge(type_, name, description)
}

/**
 ******************************************************************************
 * @file           : prometheus.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	ApiHit = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "api_hit",
		Help: "该指标记录了请求接口的次数与用时(ms),用于度量接口是否需要缓存. label { api:请求接口, method:请求方法 }",
	}, []string{"api", "method"})

	CacheHit = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "cache_hit",
		Help: "该指标记录了缓存命中,用于度量业务缓存是否需要缓存,缓存过期时间是否合理. label { key: 缓存key}",
	}, []string{"key"})

	UserBusinessErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_business_error_count",
			Help: "Count of business errors by user ID ,request uri  and error code within a time frame",
		},
		[]string{"user_id", "uri", "error_code"},
	)
)

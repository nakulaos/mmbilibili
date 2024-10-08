/**
 ******************************************************************************
 * @file           : setup.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(e *gin.Engine) {
	prometheus.MustRegister(ApiHit)
	prometheus.MustRegister(UserBusinessErrorCount)
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

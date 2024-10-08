/**
 ******************************************************************************
 * @file           : api_hit_middleware.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package middleware

import (
	"net/http"
	"time"

	"backend/pkg/prometheus"
)

func ApiHitRecord(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 开始时间
		startTime := time.Now()

		// Create a custom response writer to capture the status code
		crw := &customResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(crw, r)

		if crw.statusCode != http.StatusOK {
			return
		}

		if r.URL.Path == "/metrics" {
			return
		}

		// 只记录合法操作
		prometheus.ApiHit.With(map[string]string{
			"api":    r.URL.Path,
			"method": r.Method,
		}).Observe(float64(time.Since(startTime).Milliseconds()))
	}
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

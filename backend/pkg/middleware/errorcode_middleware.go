/**
 ******************************************************************************
 * @file           : errorcode_middleware.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package middleware

import (
	"backend/common/constant"
	"fmt"
	"net/http"

	"backend/pkg/prometheus"
)

func ErrorCodeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler
		next(w, r)

		// Check if the user ID is present in the context
		if idAny, ok := r.Context().Value(constant.UID).(string); ok {
			// Check if the error code is present in the context
			if code, ok := r.Context().Value(constant.ErrorCode).(int); ok {
				prometheus.UserBusinessErrorCount.WithLabelValues(idAny, r.URL.Path, fmt.Sprintf("%d", code)).Inc()
			}
		} else {
			// Check if the error code is present in the context
			if code, ok := r.Context().Value(constant.ErrorCode).(int); ok {
				prometheus.UserBusinessErrorCount.WithLabelValues(fmt.Sprintf("%d", constant.UnKnowUser), r.URL.Path, fmt.Sprintf("%d", code)).Inc()
			}
		}
	}
}

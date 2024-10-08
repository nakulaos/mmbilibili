package middleware

import (
	"backend/pkg/middleware"
	"net/http"
)

type ApiHitMiddleware struct {
}

func NewApiHitMiddleware() *ApiHitMiddleware {
	return &ApiHitMiddleware{}
}

func (m *ApiHitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.ApiHitRecord(next)
}

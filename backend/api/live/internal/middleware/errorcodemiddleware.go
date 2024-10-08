package middleware

import (
	"backend/pkg/middleware"
	"net/http"
)

type ErrorcodeMiddleware struct {
}

func NewErrorcodeMiddleware() *ErrorcodeMiddleware {
	return &ErrorcodeMiddleware{}
}

func (m *ErrorcodeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.ErrorCodeMiddleware(next)
}

package middleware

import (
	"backend/common/constant"
	"backend/pkg/base"
	"backend/pkg/xerror"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type AuthMiddleware struct {
	redis *redis.Redis
}

func NewAuthMiddleware(redis *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		redis: redis,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		s, err := m.redis.Get(constant.TokenBlackList + token)
		if err != nil {
			base.HttpResult(r, w, nil, errors.Wrapf(xerror.ServerError, "[AuthMiddleware] redis get error : %v", err))
			return
		}
		if s == "1" {
			base.HttpResult(r, w, nil, errors.Wrapf(xerror.AuthorizationError, "[AuthMiddleware] token is in black list : %s", token))
			return
		}
		next(w, r)
	}
}

package middleware

import (
	"github.com/redis/go-redis/v9"
	"net/http"
)

var redisClient *redis.Client

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

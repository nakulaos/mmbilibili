package dal

import (
	"backend/app/common/constant"
	"backend/app/common/ecode"
	"backend/app/http/user/conf"
	"backend/library/metric"
	"backend/library/tools"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func JWTMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var (
			token  string
			secret string
			kind   string
		)

		cf := conf.GetConf()
		token = ctx.Request.Header.Get("Authorization")

		// 验证黑名单
		result, err := RedisClient.Get(c, fmt.Sprintf("%s:%s", constant.UserTokenBlackListKey, token)).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			metric.IncrGauge(metric.LibClient, constant.PromRedisUserTokenBlackList)
			tools.SendErrResponse(c, ctx, http.StatusInternalServerError, ecode.ServerError)
			ctx.Abort()
		}
		if result != "" {
			tools.SendErrResponse(c, ctx, http.StatusUnauthorized, ecode.AuthorizationError)
			ctx.Abort()
		}

		kind = ctx.Request.Header.Get("Kind")
		if kind == "" {
			kind = "access"
		}

		if token == "" {
			tools.SendErrResponse(c, ctx, http.StatusUnauthorized, ecode.AuthorizationError)
			ctx.Abort()
		}

		if kind == "access" {
			secret = cf.App.AccessTokenSecret
		} else {
			secret = cf.App.RefreshTokenSecret
		}

		claims, err := tools.VerifyToken(token, secret)
		if err != nil {
			tools.SendErrResponse(c, ctx, http.StatusUnauthorized, ecode.AuthorizationError)
			ctx.Abort()
		}
		ctx.Set("claims", claims)
		ctx.Set("uid", claims.UserID)
		ctx.Set("username", claims.Username)

	}
}

package dal

import (
	"backend/app/rpc/user/kitex_gen/user/userrpcservice"
	"github.com/redis/go-redis/v9"
)

var (
	GlobalUserRpcClient userrpcservice.Client
	RedisClient         redis.UniversalClient
)

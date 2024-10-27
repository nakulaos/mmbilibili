package global

import (
	"backend/app/rpc/file/kitex_gen/file/fileservice"
	"backend/app/rpc/user/kitex_gen/user/userrpcservice"
	"github.com/redis/go-redis/v9"
)

var (
	UserRpcClient userrpcservice.Client
	RedisClient   redis.UniversalClient
	FileRpcClient fileservice.Client
)

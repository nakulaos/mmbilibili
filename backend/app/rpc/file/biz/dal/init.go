package dal

import (
	"backend/app/rpc/file/biz/dal/mysql"
	"backend/app/rpc/file/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

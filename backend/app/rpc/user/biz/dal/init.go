package dal

import (
	"backend/app/rpc/user/biz/dal/mysql"
	"backend/app/rpc/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

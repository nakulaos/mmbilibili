package dal

import (
	"backend/app/rpc/user/conf"
	"backend/library/initializer"
)

var (
	UserDalInstance UserDal
)

func Init(c conf.Config) {
	redisClient := initializer.InitRedisUniversal(
		c.Redis.Addrs,
		c.Redis.ClientName,
		c.Redis.DialTimeout,
		c.Redis.ReadTimeout,
		c.Redis.WriteTimeout,
		c.Redis.MaxActiveCoons,
		c.Redis.MinIdleCoons)
	cache := initializer.InitJETCache(c.UserCache, redisClient)
	db := initializer.InitGormDB(c.Mysql.DSN)
	UserDalInstance = NewUserDalImpl(cache, db, redisClient)
}

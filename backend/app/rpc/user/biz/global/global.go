package global

import (
	"backend/app/rpc/user/biz/dal"
	"backend/app/rpc/user/conf"
	"backend/library/initializer"
)

var (
	UserDal dal.UserDal
	Config  *conf.Config
)

func MustInitGlobalVal() {
	c := MustInitConf()
	redisClient := initializer.InitRedisUniversal(
		c.Redis.Addrs,
		c.Redis.ClientName,
		c.Redis.DialTimeout,
		c.Redis.ReadTimeout,
		c.Redis.WriteTimeout,
		c.Redis.MaxActiveCoons,
		c.Redis.MinIdleCoons)
	cache := initializer.InitJETCache(c.UserCache, redisClient)
	db := initializer.InitGormDBFromMysql(c.Mysql)
	UserDal = dal.NewUserDalImpl(cache, db, redisClient)
	Config = c
}

package conf

type JETCache struct {
	Name string `yaml:"name"`
	//RedisConf                  Redis  `yaml:"redis_conf"`
	LocalLFUSize               int    `yaml:"local_lfu_size"`                 // 本地缓存大小,MB
	LocalLFUTTL                int    `yaml:"local_lfu_ttl"`                  // 本地缓存过期时间,分钟
	RefreshDuration            int    `yaml:"refresh_duration"`               // 远程刷新,分钟
	StopRefreshAfterLastAccess int    `yaml:"stop_refresh_after_last_access"` // 最后一次访问后停止刷新时间,分钟
	NotFoundExpiry             int    `yaml:"not_found_expiry"`               // 未找到缓存的占位的过期时间,分钟
	RemoteExpiry               int    `yaml:"remote_expiry"`                  // 远程缓存过期时间,分钟
	SyncLocalCacheName         string `yaml:"sync_local_cache_name"`          // 同步本地缓存名称
}

type DBWithJETCache struct {
	DB       Mysql    `yaml:"db"`
	JETCache JETCache `yaml:"jet_cache"`
}

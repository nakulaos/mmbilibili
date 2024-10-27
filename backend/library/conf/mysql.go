package conf

type Mysql struct {
	DSN             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`     // 最大打开连接数
	MaxIdleConns    int    `yaml:"max_idle_conns"`     // 最大空闲连接数
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"` // 连接最大空闲时间,min
}

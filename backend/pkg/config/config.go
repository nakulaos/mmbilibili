package config

type DB struct {
	DataSource string
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type KqPusher struct {
	Brokers []string
	Topic   string
}

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type App struct {
	Salt          string
	RecommendUrl  string
	LivePusherUrl string
	HttpPusherUrl string
	FirstUsed     bool
	AppName       string
}

package conf

import (
	globalConf "backend/library/conf"
)

type Config struct {
	Env                     string
	Kitex                   globalConf.Kitex               `yaml:"kitex"`
	Registry                globalConf.Registry            `yaml:"registry"`
	OpenTelemetry           globalConf.OpenTelemetryConfig `yaml:"open_telemetry"`
	Mysql                   globalConf.Mysql               `yaml:"mysql"`
	Redis                   globalConf.Redis               `yaml:"redis"`
	UserCache               globalConf.JETCache            `yaml:"user_cache"`
	App                     App                            `yaml:"app"`
	UserRelationKafkaWriter globalConf.KafkaWriter         `yaml:"user_relation_kafka_writer"`
	UserRelevantCountWriter globalConf.KafkaWriter         `yaml:"user_relevant_count_writer"`
}

type App struct {
	Salt               string `yaml:"salt" `
	AccessTokenExpire  int64  `yaml:"access_token_expire"`
	RefreshTokenExpire int64  `yaml:"refresh_token_expire"`
	AccessTokenSecret  string `yaml:"access_token_secret"`
	RefreshTokenSecret string `yaml:"refresh_token_secret"`
	FollowingMaxCount  int64  `yaml:"following_max_count"`
	FollowingExpire    int64  `yaml:"following_expire"`
}

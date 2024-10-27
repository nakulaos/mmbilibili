package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	globalConf "backend/library/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env           string
	Hertz         globalConf.Hertz               `yaml:"hertz"`
	Registry      globalConf.Registry            `yaml:"registry"`
	Resolve       globalConf.Resolve             `yaml:"resolve"`
	OpenTelemetry globalConf.OpenTelemetryConfig `yaml:"open_telemetry"`
	UserRpc       globalConf.RPCConfig           `yaml:"user_rpc"`
	FileRpc       globalConf.RPCConfig           `yaml:"file_rpc"`
	App           App                            `yaml:"app"`
	Redis         globalConf.Redis               `yaml:"redis"`
}

type App struct {
	AccessTokenExpire  int64  `yaml:"access_token_expire"`
	RefreshTokenExpire int64  `yaml:"refresh_token_expire"`
	AccessTokenSecret  string `yaml:"access_token_secret"`
	RefreshTokenSecret string `yaml:"refresh_token_secret"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

	pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.Log.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}

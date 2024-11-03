package conf

import (
	globalConf "backend/library/conf"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config 配置结构体
type Config struct {
	UserRelevantCountService UserRelevantCountService `yaml:"user_relevant_count_service"`
}

type UserRelevantCountService struct {
	UserRelevantCountConsumer     globalConf.KafkaReader  `yaml:"user_relevant_count_consumer"`
	UserRelevantCountMysql        globalConf.Mysql        `yaml:"user_relevant_count_mysql"`
	UserRelevantCountBulkExecutor globalConf.BulkExecutor `yaml:"user_relevant_count_bulk_executor"`
}

// LoadConfig 从指定的路径加载配置文件
func LoadConfig(env string) (*Config, error) {
	configPath := fmt.Sprintf("conf/%s/conf.yaml", env)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	return &config, nil
}

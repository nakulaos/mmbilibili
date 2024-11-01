package conf

type KafkaWriter struct {
	Brokers              []string `yaml:"brokers"`
	Topic                string   `yaml:"topic"`
	EnableLocalOrder     bool     `yaml:"enable_local_order"`      // 是否开启局部顺序消息
	MaxAttempts          int      `yaml:"max_attempts"`            // 最大重试次数 次
	BatchSize            int      `yaml:"batch_size"`              // 批量发送大小 次
	BatchTimeout         int      `yaml:"batch_timeout"`           // 批量发送超时时间 ms
	BatchBytes           int      `yaml:"batch_bytes"`             // 批量发送字节数  字节数
	ReadTimeout          int      `yaml:"read_timeout"`            // 读取超时时间 ms
	WriteTimeout         int      `yaml:"write_timeout"`           // 写入超时时间 ms
	EnableAsync          bool     `yaml:"enable_async"`            // 是否异步发送
	RequiredAcks         int      `yaml:"required_acks"`           // 0: 不等待broker确认，1: 等待leader确认，-1: 等待所有副本确认
	AllowAutoCreateTopic bool     `yaml:"allow_auto_create_topic"` // 是否允许自动创建topic
}

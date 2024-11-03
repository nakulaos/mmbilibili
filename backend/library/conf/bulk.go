package conf

type BulkExecutor struct {
	TaskSize      int `yaml:"task_size"`
	FlushInterval int `yaml:"flush_interval"` // ms
}

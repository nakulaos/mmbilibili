package conf

import (
	globalConf "backend/library/conf"
)

type Config struct {
	Env           string
	Kitex         globalConf.Kitex               `yaml:"kitex"`
	Registry      globalConf.Registry            `yaml:"registry"`
	OpenTelemetry globalConf.OpenTelemetryConfig `yaml:"open_telemetry"`
	Mysql         globalConf.Mysql               `yaml:"mysql"`
	MinIO         globalConf.MinIO               `yaml:"minio"`
	App           App                            `yaml:"app"`
}

type App struct {
	FileMaxSize                      int64 `yaml:"file_max_size"`                         //mb，单文件最大大小
	ChunkMaxSize                     int64 `yaml:"chunk_max_size"`                        //mb，单块最大大小
	PresignedUploadPartUrlExpireTime int64 `yaml:"presigned_upload_part_url_expire_time"` // 分块上传url过期时间,min
}

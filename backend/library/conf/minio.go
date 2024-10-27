package conf

type MinIO struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKey       string `yaml:"access_key"`
	SecretAccessKey string `yaml:"secret_access_key"`
	BucketName      string `yaml:"bucket_name"`
	UseSSL          bool   `yaml:"use_ssl"`
	BasePath        string `yaml:"base_path"`
	Location        string `yaml:"location"`
}

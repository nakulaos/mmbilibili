package initializer

import (
	globalConf "backend/library/conf"
	"backend/library/minio_ext"
	"github.com/minio/minio-go/v6"
	"log"
)

func InitMinIOClient(c *globalConf.MinIO) *minio.Client {
	minioClient, err := minio.New(c.Endpoint, c.AccessKey, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		log.Fatal(err)
	}
	return minioClient
}

func InitMinIOCoreClient(c *globalConf.MinIO) *minio.Core {
	minioCore, err := minio.NewCore(c.Endpoint, c.AccessKey, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		log.Fatal(err)
	}
	return minioCore
}

func InitMinIOExtClient(c *globalConf.MinIO) *minio_ext.Client {
	minioClient, err := minio_ext.New(c.Endpoint, c.AccessKey, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		log.Fatal(err)
	}
	return minioClient
}

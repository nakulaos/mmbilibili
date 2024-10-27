package global

import (
	"backend/app/rpc/file/biz/dal"
	"backend/app/rpc/file/conf"
	"backend/library/initializer"
	"backend/library/minio_ext"
	"github.com/minio/minio-go/v6"
)

var (
	Config          *conf.Config
	MinioClient     *minio.Client
	MinioCoreClient *minio.Core
	MinioExtClient  *minio_ext.Client
	FileDal         dal.FileDal
)

func MustInitGlobalVal() {
	c := MustInitConf()
	Config = c
	MinioClient = initializer.InitMinIOClient(&c.MinIO)
	MinioCoreClient = initializer.InitMinIOCoreClient(&c.MinIO)
	MinioExtClient = initializer.InitMinIOExtClient(&c.MinIO)
	FileDal = dal.NewFileDalImpl(initializer.InitGormDBFromMysql(c.Mysql))
}

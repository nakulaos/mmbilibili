package initializer

import (
	globalConf "backend/library/conf"
	"context"
	"fmt"
	"github.com/minio/minio-go/v6"
	"reflect"
	"testing"
)

func TestInitMinIOClient(t *testing.T) {
	type args struct {
		c *globalConf.MinIO
	}
	tests := []struct {
		name string
		args args
		want *minio.Client
	}{
		// TODO: Add test cases.
		{
			args: args{c: &globalConf.MinIO{
				Endpoint:        "localhost:9003",
				AccessKey:       "Yz5N8GaP07eIu2vOpOre",
				SecretAccessKey: "qGg2D5AqyNW49jcr6jW2m1SB82TfZnEwyAYsCnQU",
				BucketName:      "test",
				UseSSL:          false,
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InitMinIOClient(tt.args.c)
			info, err := got.ListBuckets(context.Background())
			if err != nil {
				t.Errorf("get list buckets error: %v", err)
			}
			fmt.Sprintf("info: %v", info)
		})
	}
}

func TestInitMinIOCoreClient(t *testing.T) {
	type args struct {
		c *globalConf.MinIO
	}
	tests := []struct {
		name string
		args args
		want *minio.Core
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitMinIOCoreClient(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMinIOCoreClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

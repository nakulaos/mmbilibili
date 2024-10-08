package tools

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

// PutObject now := time.Now().Format("20060102150405")
// key := fmt.Sprintf("%s/%d__%s__%s", _avatarPrefix_, req.User.ID, now, fileName)
func PutObject(accessKey, secretKey, bucket, objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool, keyToOverWrite bool) (string, error) {
	// 创建一个上传凭证
	var putPolicy storage.PutPolicy
	if !keyToOverWrite {
		putPolicy = storage.PutPolicy{
			Scope: bucket,
		}
	} else {
		putPolicy = storage.PutPolicy{
			Scope: fmt.Sprintf("%s:%s", bucket, objectKey),
		}
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 是否使用https域名进行上传
	cfg.UseHTTPS = false
	// 是否使用CDN上传加速
	cfg.UseCdnDomains = false

	uploader := storage.NewFormUploader(&cfg)

	// PutRet 为七牛标准的上传回复内容。
	// 如果使用了上传回调或者自定义了returnBody，那么需要根据实际情况，自己自定义一个返回值结构体
	ret := storage.PutRet{}

	// PutExtra 为表单上传的额外可选项
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}

	err := uploader.Put(
		context.Background(),
		&ret,
		upToken,
		objectKey,
		reader,
		objectSize,
		&putExtra,
	)

	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func PutVideo(accessKey, secretKey, bucket, objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool, keyToOverWrite bool) (string, error) {
	// 创建一个上传凭证
	var putPolicy storage.PutPolicy
	if !keyToOverWrite {
		putPolicy = storage.PutPolicy{
			Scope: bucket,
		}
	} else {
		putPolicy = storage.PutPolicy{
			Scope: fmt.Sprintf("%s:%s", bucket, objectKey),
		}
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 是否使用https域名进行上传
	cfg.UseHTTPS = false
	// 是否使用CDN上传加速
	cfg.UseCdnDomains = false

	uploader := storage.NewFormUploader(&cfg)

	// PutRet 为七牛标准的上传回复内容。
	// 如果使用了上传回调或者自定义了returnBody，那么需要根据实际情况，自己自定义一个返回值结构体
	ret := storage.PutRet{}

	// PutExtra 为表单上传的额外可选项
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}

	err := uploader.Put(
		context.Background(),
		&ret,
		upToken,
		objectKey,
		reader,
		objectSize,
		&putExtra,
	)

	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

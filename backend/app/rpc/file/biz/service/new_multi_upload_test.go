package service

import (
	"backend/app/rpc/file/biz/global"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"testing"
)

func TestNewMultiUpload_Run(t *testing.T) {
	global.MustInitGlobalVal()
	ctx := context.Background()
	s := NewNewMultiUploadService(ctx)
	// init req and assert value

	req := &file.NewMultiUploadReq{
		FileHash:         "a2b513d9c79dde5e2992977b35ed7b9a",
		ChunkTotalNumber: 3,
		ChunkSize:        3,
		FileName:         "test.txt",
		UserID:           1,
		FileSize:         1024 * 3,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

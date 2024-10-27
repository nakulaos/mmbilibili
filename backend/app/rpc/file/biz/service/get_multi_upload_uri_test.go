package service

import (
	"backend/app/rpc/file/biz/global"
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"testing"
)

func TestGetMultiUploadUri_Run(t *testing.T) {
	global.MustInitGlobalVal()
	ctx := context.Background()
	s := NewGetMultiUploadUriService(ctx)
	// init req and assert value

	req := &file.GetMultiUploadUriReq{
		UUID:             "989f4ed0-bc7f-48c2-9461-f070532ab279",
		UploadID:         "NDEzMzY3NGUtZjNhMy00NmFkLTkwZTEtN2M2NmE5NWYzZjI1LmI3MzYwMGYwLWJhNzYtNGQ3Yy1iYjM1LTg2YjQxYjgzNjVhNA",
		ChunkTotalNumber: 3,
		UserID:           1,
		ChunkID:          1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

package service

import (
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"testing"
)

func TestGetSuccessChunks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetSuccessChunksService(ctx)
	// init req and assert value

	req := &file.GetSuccessChunksReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

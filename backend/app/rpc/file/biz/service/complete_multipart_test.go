package service

import (
	file "backend/app/rpc/file/kitex_gen/file"
	"context"
	"testing"
)

func TestCompleteMultipart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCompleteMultipartService(ctx)
	// init req and assert value

	req := &file.CompleteMultipartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

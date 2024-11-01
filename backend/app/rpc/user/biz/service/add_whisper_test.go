package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"testing"
)

func TestAddWhisper_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddWhisperService(ctx)
	// init req and assert value

	req := &user.AddWhisperReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

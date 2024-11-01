package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"testing"
)

func TestDelBlack_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDelBlackService(ctx)
	// init req and assert value

	req := &user.DelBlackReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

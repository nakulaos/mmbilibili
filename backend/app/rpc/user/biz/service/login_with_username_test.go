package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"testing"
)

func TestLoginWithUsername_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginWithUsernameService(ctx)
	// init req and assert value

	req := &user.LoginWithUsernameReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

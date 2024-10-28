package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"testing"
)

func TestRefreshToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefreshTokenService(ctx)
	// init req and assert value

	req := &user.RefreshTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

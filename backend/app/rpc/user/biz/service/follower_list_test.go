package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"testing"
)

func TestFollowerList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFollowerListService(ctx)
	// init req and assert value

	req := &user.FollowerListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

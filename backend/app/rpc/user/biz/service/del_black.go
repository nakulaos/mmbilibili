package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type DelBlackService struct {
	ctx context.Context
} // NewDelBlackService new DelBlackService
func NewDelBlackService(ctx context.Context) *DelBlackService {
	return &DelBlackService{ctx: ctx}
}

// Run create note info
func (s *DelBlackService) Run(req *user.DelBlackReq) (resp *user.DelBlackResp, err error) {
	// Finish your business logic.

	return
}

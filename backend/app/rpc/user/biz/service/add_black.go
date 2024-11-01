package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type AddBlackService struct {
	ctx context.Context
} // NewAddBlackService new AddBlackService
func NewAddBlackService(ctx context.Context) *AddBlackService {
	return &AddBlackService{ctx: ctx}
}

// Run create note info
func (s *AddBlackService) Run(req *user.AddBlackReq) (resp *user.AddBlackResp, err error) {
	// Finish your business logic.

	return
}

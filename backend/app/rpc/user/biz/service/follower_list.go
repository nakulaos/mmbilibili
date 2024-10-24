package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type FollowerListService struct {
	ctx context.Context
} // NewFollowerListService new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

// Run create note info
func (s *FollowerListService) Run(req *user.FollowerListReq) (resp *user.FollowerListResp, err error) {
	// Finish your business logic.

	return
}

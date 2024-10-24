package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type FollowUserService struct {
	ctx context.Context
} // NewFollowUserService new FollowUserService
func NewFollowUserService(ctx context.Context) *FollowUserService {
	return &FollowUserService{ctx: ctx}
}

// Run create note info
func (s *FollowUserService) Run(req *user.FollowUserReq) (resp *user.FollowUserResp, err error) {
	// Finish your business logic.

	return
}

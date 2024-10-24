package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type FollowingListService struct {
	ctx context.Context
} // NewFollowingListService new FollowingListService
func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{ctx: ctx}
}

// Run create note info
func (s *FollowingListService) Run(req *user.FollowingListReq) (resp *user.FollowingListResp, err error) {
	// Finish your business logic.

	return
}

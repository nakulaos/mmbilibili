package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type FriendListService struct {
	ctx context.Context
} // NewFriendListService new FriendListService
func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

// Run create note info
func (s *FriendListService) Run(req *user.FriendListReq) (resp *user.FriendListResp, err error) {
	// Finish your business logic.

	return
}

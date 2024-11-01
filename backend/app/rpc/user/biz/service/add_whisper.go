package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type AddWhisperService struct {
	ctx context.Context
} // NewAddWhisperService new AddWhisperService
func NewAddWhisperService(ctx context.Context) *AddWhisperService {
	return &AddWhisperService{ctx: ctx}
}

// Run create note info
func (s *AddWhisperService) Run(req *user.AddWhisperReq) (resp *user.AddWhisperResp, err error) {
	// Finish your business logic.

	return
}

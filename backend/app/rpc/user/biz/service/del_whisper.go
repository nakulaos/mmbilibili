package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type DelWhisperService struct {
	ctx context.Context
} // NewDelWhisperService new DelWhisperService
func NewDelWhisperService(ctx context.Context) *DelWhisperService {
	return &DelWhisperService{ctx: ctx}
}

// Run create note info
func (s *DelWhisperService) Run(req *user.DelWhisperReq) (resp *user.DelWhisperResp, err error) {
	// Finish your business logic.

	return
}

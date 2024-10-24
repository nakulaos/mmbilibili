package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type LoginWithEmailService struct {
	ctx context.Context
} // NewLoginWithEmailService new LoginWithEmailService
func NewLoginWithEmailService(ctx context.Context) *LoginWithEmailService {
	return &LoginWithEmailService{ctx: ctx}
}

// Run create note info
func (s *LoginWithEmailService) Run(req *user.LoginWithEmailReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	return
}

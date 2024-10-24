package service

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type LoginWithPhoneService struct {
	ctx context.Context
} // NewLoginWithPhoneService new LoginWithPhoneService
func NewLoginWithPhoneService(ctx context.Context) *LoginWithPhoneService {
	return &LoginWithPhoneService{ctx: ctx}
}

// Run create note info
func (s *LoginWithPhoneService) Run(req *user.LoginWithPhoneReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	return
}

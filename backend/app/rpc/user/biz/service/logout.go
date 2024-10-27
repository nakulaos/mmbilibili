package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/global"
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	accessToken := req.AccessToken
	refreshToken := req.RefreshToken
	if err = global.UserDal.AddTokenToBlackList(s.ctx, accessToken); err != nil {
		return nil, ecode.ServerError
	}
	if err = global.UserDal.AddTokenToBlackList(s.ctx, refreshToken); err != nil {
		return nil, ecode.ServerError
	}

	return
}

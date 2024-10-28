package service

import (
	"backend/app/common/constant"
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/global"
	"backend/app/rpc/user/biz/model"
	user "backend/app/rpc/user/kitex_gen/user"
	"backend/library/metric"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RefreshTokenService struct {
	ctx context.Context
} // NewRefreshTokenService new RefreshTokenService
func NewRefreshTokenService(ctx context.Context) *RefreshTokenService {
	return &RefreshTokenService{ctx: ctx}
}

// Run create note info
func (s *RefreshTokenService) Run(req *user.RefreshTokenReq) (resp *user.RefreshTokenResp, err error) {
	var (
		refreshToken    = req.RefreshToken
		newAccessToken  string
		newRefreshToken string
		uid             = req.UserId
		userModel       *model.User
	)

	if err = global.UserDal.AddTokenToBlackList(s.ctx, refreshToken); err != nil {
		return nil, ecode.ServerError
	}

	if userModel, err = global.UserDal.GetUserByID(s.ctx, uid); err != nil {
		return nil, ecode.ServerError
	}

	var (
		accessSecret  = global.Config.App.AccessTokenSecret
		refreshSecret = global.Config.App.RefreshTokenSecret
		accessExpire  = global.Config.App.AccessTokenExpire
		refreshExpire = global.Config.App.RefreshTokenExpire
	)

	newAccessToken, newRefreshToken, err = tools.GenerateDoubleToken(userModel.ID, userModel.Username, accessSecret, refreshSecret, accessExpire, refreshExpire)
	if err != nil {
		metric.IncrGauge(metric.BusinessError, constant.PromGenerateTokenError)
		klog.Errorf("registerService.tools.generateDoubleToken(%d,%s) error:%v", userModel.ID, userModel.Username, err)
		return nil, nil
	}

	resp = &user.RefreshTokenResp{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}

	return resp, nil
}

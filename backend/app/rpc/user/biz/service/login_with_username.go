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
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type LoginWithUsernameService struct {
	ctx context.Context
} // NewLoginWithUsernameService new LoginWithUsernameService
func NewLoginWithUsernameService(ctx context.Context) *LoginWithUsernameService {
	return &LoginWithUsernameService{ctx: ctx}
}

// Run create note info
func (s *LoginWithUsernameService) Run(req *user.LoginWithUsernameReq) (resp *user.LoginResp, err error) {
	/*
		1. 判断用户名和密码是否为空
		2. 判断用户是否存在
		3. 密码校对
		4. 判断用户是否被禁用
		5. 获取用户相关统计信息
		6. 生成双token
		7. 返回用户信息
	*/

	username := req.Username
	password := req.Password

	if username == "" || password == "" {
		return nil, ecode.InvalidParamsError.WithTemplateData(map[string]string{"Params": "username or password is empty"})
	}

	userModel, err := global.UserDal.GetUserByUserName(s.ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.UserNotExistError.WithTemplateData(map[string]string{"UID": username})
		}
		return nil, ecode.ServerError
	}

	cryptPassword := tools.PasswordEncrypt(userModel.Salt, global.Config.App.Salt, password)
	if userModel.Password != cryptPassword {
		return nil, ecode.PassWordError
	}

	if userModel.Status == model.AttrBanLoginState {
		return nil, ecode.UserDisableError
	}

	userStats := &model.UserRelevantCount{}
	userStats, err = global.UserDal.GetUserRelevantCountByID(s.ctx, userModel.ID)
	if err != nil {
		return nil, ecode.ServerError
	}

	var (
		accessSecret  = global.Config.App.AccessTokenSecret
		refreshSecret = global.Config.App.RefreshTokenSecret
		accessExpire  = global.Config.App.AccessTokenExpire
		refreshExpire = global.Config.App.RefreshTokenExpire
	)

	accessToken, refreshToken, err := tools.GenerateDoubleToken(userModel.ID, userModel.Username, accessSecret, refreshSecret, accessExpire, refreshExpire)
	if err != nil {
		metric.IncrGauge(metric.BusinessError, constant.PromGenerateTokenError)
		klog.Errorf("registerService.tools.generateDoubleToken(%d,%s) error:%v", userModel.ID, userModel.Username, err)
		return nil, nil
	}

	resp = &user.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       userModel.ID,
		UserInfo: &user.User{
			Id:             userModel.ID,
			Username:       userModel.Username,
			Role:           uint32(userModel.Role),
			Avatar:         userModel.Avatar,
			Nickname:       userModel.Nickname,
			Gender:         uint32(userModel.Gender),
			Phone:          tools.MaskPhone(userModel.Phone),
			Email:          tools.MaskEmail(userModel.Email),
			FollowerCount:  userStats.FollowerCount,
			FollowingCount: userStats.FollowingCount,
			LikeCount:      userStats.LikeCount,
			StarCount:      userStats.StarCount,
			SelfStarCount:  userStats.SelfStarCount,
			SelfLikeCount:  userStats.SelfLikeCount,
			LiveCount:      userStats.LiveCount,
			WorkCount:      userStats.WorkCount,
			FriendCount:    userStats.FriendCount,
		},
	}
	return resp, nil
}

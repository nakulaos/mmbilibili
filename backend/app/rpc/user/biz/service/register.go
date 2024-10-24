package service

import (
	"backend/app/common/constant"
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/dal"
	"backend/app/rpc/user/biz/model"
	"backend/app/rpc/user/conf"
	user "backend/app/rpc/user/kitex_gen/user"
	"backend/library/metric"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.LoginResp, err error) {
	/*
		1. 判断用户名和密码是否为空
		2. 判断用户是否存在
		3. 生成用户盐值
		4. 创建用户
		5. 生成双token
		6. mask 邮箱和手机号
		7. 返回用户信息
	*/

	username := req.Username
	password := req.Password

	if username == "" || password == "" {
		return nil, ecode.InvalidParamsError.WithTemplateData(map[string]string{"Params": "username or password is empty"})
	}

	f, err := dal.UserDalInstance.ExistUserByUserName(s.ctx, username)
	if err != nil {
		return nil, ecode.ServerError
	}
	if f {
		metric.IncrGauge(metric.BusinessError, constant.PromUserExistError)
		return nil, ecode.UserExistError.WithTemplateData(map[string]string{"UID": username})
	}

	salt, _ := tools.GenerateSalt()
	newUser := &model.User{
		Username: username,
		Password: tools.PasswordEncrypt(salt, conf.GetConf().App.Salt, password),
		Salt:     salt,
	}
	err = dal.UserDalInstance.CreateUser(s.ctx, newUser)

	var (
		accessSecret  = conf.GetConf().App.AccessTokenSecret
		refreshSecret = conf.GetConf().App.RefreshTokenSecret
		accessExpire  = conf.GetConf().App.AccessTokenExpire
		refreshExpire = conf.GetConf().App.RefreshTokenExpire
	)

	accessToken, refreshToken, err := tools.GenerateDoubleToken(newUser.ID, newUser.Username, accessSecret, refreshSecret, accessExpire, refreshExpire)
	if err != nil {
		metric.IncrGauge(metric.BusinessError, constant.PromGenerateTokenError)
		klog.Errorf("registerService.tools.generateDoubleToken(%d,%s) error:%v", newUser.ID, newUser.Username, err)
		return nil, nil
	}

	resp = &user.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       newUser.ID,
		UserInfo: &user.User{
			Id:       newUser.ID,
			Username: newUser.Username,
			Nickname: newUser.Nickname,
			Avatar:   newUser.Avatar,
			Gender:   uint32(newUser.Gender),
			Role:     uint32(1),
			Phone:    tools.MaskPhone(newUser.Phone),
			Email:    tools.MaskEmail(newUser.Email),
		},
	}

	return resp, nil
}

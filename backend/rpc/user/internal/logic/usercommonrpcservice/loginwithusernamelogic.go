package usercommonrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/modelcache"
	"backend/dao/statistics"
	"backend/pkg/crypt"
	"backend/pkg/jwt"
	"backend/pkg/xerror"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithUsernameLogic {
	return &LoginWithUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithUsernameLogic) LoginWithUsername(in *user.LoginWithUsernameReq) (*user.LoginResp, error) {
	var err error
	userModel, _ := modelcache.GetUserModelCacheFromUsername(l.svcCtx.Cache, l.ctx, in.Username)
	if userModel.ID == 0 {
		// 从数据库中获取
		userModel, err = l.svcCtx.Dao.User.WithContext(l.ctx).Where(l.svcCtx.Dao.User.Username.Eq(in.Username)).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerror.UserNotExistError.WithTemplateData(map[string]interface{}{
					"UID": in.Username,
				}), "[LoginWithUsername] get user failed: %v,because the user is not exist", err)
			} else {
				return nil, errors.Wrapf(xerror.ServerError, "[LoginWithUsername] get user failed: %v", err)
			}
		}
	}

	password := crypt.PasswordEncrypt(userModel.Salt, l.svcCtx.Config.App.Salt, in.Password)
	if userModel.Password != password {
		return nil, errors.Wrapf(xerror.PassWordError, "[LoginWithUsername] username :%s", in.Username)
	}

	if userModel.Status != constant.UserStatusNormal {
		return nil, errors.Wrapf(xerror.UserDisableError, "[LoginWithUsername] username :%s", in.Username)
	}

	token, err := jwt.MakeToken(int(userModel.ID), int(userModel.Role), l.svcCtx.Config.Jwt)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LoginWithUsername] username :%s make token fail", in.Username)
	}

	// 获取用户统计数据
	userStats, err := statistics.GetUserRelevantCount(l.svcCtx.Redis, l.ctx, int(userModel.ID))
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[UpdateUserInfoLogic] GetUserRelevantCount error:%v", err)
	}

	var resp user.LoginResp
	resp.AccessToken = token
	resp.UserId = int64(userModel.ID)
	resp.UserInfo = &user.User{
		Id:             uint32(userModel.ID),
		Username:       userModel.Username,
		Role:           uint32(userModel.Role),
		Avatar:         userModel.Avatar,
		Nickname:       userModel.Nickname,
		Gender:         uint32(userModel.Gender),
		Phone:          userModel.Phone,
		Email:          userModel.Email,
		FollowerCount:  int32(userStats.FollowerCount),
		FollowingCount: int32(userStats.FollowingCount),
		LikeCount:      int32(userStats.LikeCount),
		StarCount:      int32(userStats.StarCount),
		SelfStarCount:  int32(userStats.SelfStarCount),
		SelfLikeCount:  int32(userStats.SelfLikeCount),
		LiveCount:      int32(userStats.LiveCount),
		WorkCount:      int32(userStats.WorkCount),
		FriendCount:    int32(userStats.FriendCount),
	}

	go func() {
		// 缓存用户信息
		modelcache.SetUserModelCache(l.svcCtx.Cache, userModel)
	}()

	return &resp, nil
}

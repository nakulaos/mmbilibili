package usercommonrpcservicelogic

import (
	"backend/common/constant"
	"backend/pkg/xerror"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	err := l.svcCtx.Redis.SetexCtx(context.Background(), constant.TokenBlackList+in.LoginToken, strconv.Itoa(1), int(l.svcCtx.Config.Jwt.AccessExpire)*int(time.Hour))
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[Logout] add token to black list failed, token: %s, err: %+v", in.LoginToken, err)
	}
	return nil, nil
}

package usercommonrpcservicelogic

import (
	"context"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithPhoneLogic {
	return &LoginWithPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithPhoneLogic) LoginWithPhone(in *user.LoginWithPhoneReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResp{}, nil
}

package userfilerpcservicelogic

import (
	"context"

	"backend/rpc/user/internal/svc"
	"backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUploadFileLogic {
	return &UserUploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUploadFileLogic) UserUploadFile(in *user.UserUploadFileReq) (*user.UserUploadFileResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserUploadFileResp{}, nil
}

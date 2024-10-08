package file

import (
	"context"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUploadFileLogic {
	return &UserUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUploadFileLogic) UserUploadFile(req *types.UserUploadFileReq) (resp *types.UserUploadFileResp, err error) {
	// todo: add your logic here and delete this line

	return
}

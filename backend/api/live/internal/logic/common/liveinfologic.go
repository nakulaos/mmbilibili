package common

import (
	"context"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveInfoLogic {
	return &LiveInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveInfoLogic) LiveInfo(req *types.LiveDetailReq) (resp *types.LiveDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}

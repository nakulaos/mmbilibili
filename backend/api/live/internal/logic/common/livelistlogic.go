package common

import (
	"context"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveListLogic {
	return &LiveListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveListLogic) LiveList(req *types.LiveListReq) (resp *types.LiveListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

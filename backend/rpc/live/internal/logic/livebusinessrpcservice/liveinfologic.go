package livebusinessrpcservicelogic

import (
	"context"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveInfoLogic {
	return &LiveInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveInfoLogic) LiveInfo(in *live.LiveDetailReq) (*live.LiveDetailResp, error) {
	// todo: add your logic here and delete this line

	return &live.LiveDetailResp{}, nil
}

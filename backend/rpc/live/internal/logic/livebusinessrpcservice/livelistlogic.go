package livebusinessrpcservicelogic

import (
	"context"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveListLogic {
	return &LiveListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveListLogic) LiveList(in *live.LiveListReq) (*live.LiveListResp, error) {
	// todo: add your logic here and delete this line

	return &live.LiveListResp{}, nil
}

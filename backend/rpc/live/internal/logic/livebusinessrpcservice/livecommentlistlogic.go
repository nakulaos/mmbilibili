package livebusinessrpcservicelogic

import (
	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveCommentListLogic {
	return &LiveCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveCommentListLogic) LiveCommentList(in *live.LiveCommentListReq) (*live.LiveCommentListResp, error) {
	return nil, nil
}

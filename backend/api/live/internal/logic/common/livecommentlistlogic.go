package common

import (
	"context"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveCommentListLogic {
	return &LiveCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveCommentListLogic) LiveCommentList(req *types.LiveCommentListReq) (resp *types.LiveCommentListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

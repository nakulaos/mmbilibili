package callback

import (
	"backend/pkg/xerror"
	"backend/rpc/live/live"
	"context"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnStreamChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnStreamChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnStreamChangeLogic {
	return &OnStreamChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnStreamChangeLogic) OnStreamChange(req *types.OnStreamChangeReq) (resp *types.OnStreamChangeResp, err error) {
	in := &live.OnStreamChangeReq{
		App:           req.App,
		Regist:        req.Regist,
		Schema:        req.Schema,
		Stream:        req.Stream,
		Vhost:         req.VHost,
		MediaServerId: req.MediaServerID,
	}
	_, err = l.svcCtx.LiveCallbackServiceClient.OnStreamChange(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.OnStreamChangeResp{}, xerror.Ok
}

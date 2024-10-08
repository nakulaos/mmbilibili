package callback

import (
	"backend/pkg/xerror"
	"backend/rpc/live/live"
	"context"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnPublicStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnPublicStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnPublicStreamLogic {
	return &OnPublicStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnPublicStreamLogic) OnPublicStream(req *types.OnPublicStreamReq) (resp *types.OnPublicStreamResp, err error) {
	in := &live.OnPublicStreamReq{
		App:           req.App,
		Schema:        req.Schema,
		Stream:        req.Stream,
		Vhost:         req.VHost,
		MediaServerId: req.MediaServerID,
		Ip:            req.IP,
		Port:          uint32(req.Port),
		Params:        req.Params,
		Id:            req.ID,
	}

	_, err = l.svcCtx.LiveCallbackServiceClient.OnPublicStream(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &types.OnPublicStreamResp{}, xerror.Ok
}

package auth

import (
	"backend/rpc/live/live"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveLikeLogic {
	return &LiveLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveLikeLogic) LiveLike(req *types.LiveLikeReq) (*types.LiveLikeResp, error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := live.LiveLikeReq{
		UserId: uint32(uid),
		LiveId: uint32(req.LiveID),
		Action: int32(req.Action),
	}

	liveLikeResp, err := l.svcCtx.LiveBusinessServiceClient.LiveLike(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "live like error: uid: %d ,lid :%d", uid, in.LiveId)
	}
	var resp types.LiveLikeResp
	resp.LikeCount = int(liveLikeResp.LikeCount)

	return &resp, nil
}

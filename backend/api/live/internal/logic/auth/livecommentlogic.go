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

type LiveCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveCommentLogic {
	return &LiveCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveCommentLogic) LiveComment(req *types.LiveCommentReq) (resp *types.LiveCommentResp, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := live.LiveCommentReq{
		UserId:   uint32(uid),
		LiveId:   uint32(req.LiveID),
		Content:  req.Content,
		SendTime: req.SendTime,
	}
	liveCommentResp, err := l.svcCtx.LiveBusinessServiceClient.LiveComment(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "live comment error: uid: %d ,lid :%d,content:%s", uid, in.LiveId, in.Content)
	}

	resp = &types.LiveCommentResp{}
	resp.CommentID = uint(liveCommentResp.CommentId)

	return resp, nil
}

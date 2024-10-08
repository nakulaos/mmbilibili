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

type LiveDeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveDeleteCommentLogic {
	return &LiveDeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveDeleteCommentLogic) LiveDeleteComment(req *types.LiveDeleteCommentReq) (resp *types.LiveCommentResp, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := live.LiveDeleteCommentReq{
		UserId:    uint32(uid),
		CommentId: uint32(req.CommentID),
	}
	liveDeleteCommentResp, err := l.svcCtx.LiveBusinessServiceClient.LiveDeleteComment(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "live delete comment error: uid: %d ,commentId:%d", uid, in.CommentId)
	}

	resp = &types.LiveCommentResp{}
	resp.CommentID = uint(liveDeleteCommentResp.CommentId)

	return resp, nil
}

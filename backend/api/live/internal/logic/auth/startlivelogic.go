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

type StartLiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLiveLogic {
	return &StartLiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartLiveLogic) StartLive(req *types.StartLiveReq) (*types.LiveDetailResp, error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := live.StartLiveReq{
		Uid:         int32(uid),
		Title:       req.Title,
		Cover:       req.Cover,
		Category:    req.Category,
		Description: req.Description,
		Tags:        req.Tags,
		Partition:   req.Partition,
	}

	startLiveResp, err := l.svcCtx.LiveBusinessServiceClient.StartLive(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "start live error: uid: %d", uid)
	}

	var resp types.LiveDetailResp
	resp.LiveInfo = types.LiveInfo{
		LiveID:       uint(startLiveResp.LiveInfo.LiveId),
		UserID:       uint(startLiveResp.LiveInfo.UserId),
		Title:        startLiveResp.LiveInfo.Title,
		Cover:        startLiveResp.LiveInfo.Cover,
		Status:       startLiveResp.LiveInfo.Status,
		StartTime:    startLiveResp.LiveInfo.StartTime,
		EndTime:      startLiveResp.LiveInfo.EndTime,
		WatchCount:   int(startLiveResp.LiveInfo.WatchCount),
		LikeCount:    int(startLiveResp.LiveInfo.LikeCount),
		CommentCount: int(startLiveResp.LiveInfo.CommentCount),
		ShareCount:   int(startLiveResp.LiveInfo.ShareCount),
		IsLike:       startLiveResp.LiveInfo.IsLike,
		IsFollow:     startLiveResp.LiveInfo.IsFollow,
		IsStar:       startLiveResp.LiveInfo.IsStar,
		IsSelf:       startLiveResp.LiveInfo.IsSelf,
		Type:         int(startLiveResp.LiveInfo.Type),
		Description:  startLiveResp.LiveInfo.Description,
		PlayerUrl:    startLiveResp.LiveInfo.PlayerUrl,
		CoverUrl:     startLiveResp.LiveInfo.CoverUrl,
		IsOver:       startLiveResp.LiveInfo.IsOver,
		Category:     startLiveResp.LiveInfo.Category,
		Tags:         startLiveResp.LiveInfo.Tags,
		Partition:    startLiveResp.LiveInfo.Partition,
		RoomID:       uint(startLiveResp.LiveInfo.RoomId),
		Token:        startLiveResp.LiveInfo.Token,
	}
	resp.LiveInfo.Author = types.User{
		Id:             startLiveResp.LiveInfo.Author.Id,
		Nickname:       startLiveResp.LiveInfo.Author.Nickname,
		Username:       startLiveResp.LiveInfo.Author.Username,
		Avatar:         startLiveResp.LiveInfo.Author.Avatar,
		FollowerCount:  int(startLiveResp.LiveInfo.Author.FollowerCount),
		FollowingCount: int(startLiveResp.LiveInfo.Author.FollowingCount),
		LikeCount:      int(startLiveResp.LiveInfo.Author.LikeCount),
		StarCount:      int(startLiveResp.LiveInfo.Author.StarCount),
		SelfStarCount:  int(startLiveResp.LiveInfo.Author.SelfStarCount),
		SelfLikeCount:  int(startLiveResp.LiveInfo.Author.SelfLikeCount),
		LiveCount:      int(startLiveResp.LiveInfo.Author.LiveCount),
		WorkCount:      int(startLiveResp.LiveInfo.Author.WorkCount),
		FriendCount:    int(startLiveResp.LiveInfo.Author.FriendCount),
		Role:           startLiveResp.LiveInfo.Author.Role,
		Gender:         startLiveResp.LiveInfo.Author.Gender,
		Status:         uint(startLiveResp.LiveInfo.Author.Status),
	}
	return &resp, nil
}

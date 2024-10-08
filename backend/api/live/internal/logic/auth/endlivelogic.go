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

type EndLiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndLiveLogic {
	return &EndLiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EndLiveLogic) EndLive(req *types.EndLiveReq) (*types.LiveDetailResp, error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := live.EndLiveReq{
		UserId: uint32(uid),
		LiveId: uint32(req.LiveID),
	}

	endLiveResp, err := l.svcCtx.LiveBusinessServiceClient.EndLive(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "end live error: uid: %d ,lid :%d", uid, in.LiveId)
	}

	var resp types.LiveDetailResp
	resp.LiveInfo = types.LiveInfo{
		LiveID:       uint(endLiveResp.LiveInfo.LiveId),
		UserID:       uint(endLiveResp.LiveInfo.UserId),
		Title:        endLiveResp.LiveInfo.Title,
		Cover:        endLiveResp.LiveInfo.Cover,
		Status:       endLiveResp.LiveInfo.Status,
		StartTime:    endLiveResp.LiveInfo.StartTime,
		EndTime:      endLiveResp.LiveInfo.EndTime,
		WatchCount:   int(endLiveResp.LiveInfo.WatchCount),
		LikeCount:    int(endLiveResp.LiveInfo.LikeCount),
		CommentCount: int(endLiveResp.LiveInfo.CommentCount),
		ShareCount:   int(endLiveResp.LiveInfo.ShareCount),
		IsLike:       endLiveResp.LiveInfo.IsLike,
		IsFollow:     endLiveResp.LiveInfo.IsFollow,
		IsStar:       endLiveResp.LiveInfo.IsStar,
		IsSelf:       endLiveResp.LiveInfo.IsSelf,
		Type:         int(endLiveResp.LiveInfo.Type),
		Description:  endLiveResp.LiveInfo.Description,
		PlayerUrl:    endLiveResp.LiveInfo.PlayerUrl,
		CoverUrl:     endLiveResp.LiveInfo.CoverUrl,
		IsOver:       endLiveResp.LiveInfo.IsOver,
		Category:     endLiveResp.LiveInfo.Category,
		Tags:         endLiveResp.LiveInfo.Tags,
		Partition:    endLiveResp.LiveInfo.Partition,
		RoomID:       uint(endLiveResp.LiveInfo.RoomId),
		Token:        endLiveResp.LiveInfo.Token,
	}
	resp.LiveInfo.Author = types.User{
		Id:             endLiveResp.LiveInfo.Author.Id,
		Nickname:       endLiveResp.LiveInfo.Author.Nickname,
		Username:       endLiveResp.LiveInfo.Author.Username,
		Avatar:         endLiveResp.LiveInfo.Author.Avatar,
		FollowerCount:  int(endLiveResp.LiveInfo.Author.FollowerCount),
		FollowingCount: int(endLiveResp.LiveInfo.Author.FollowingCount),
		LikeCount:      int(endLiveResp.LiveInfo.Author.LikeCount),
		StarCount:      int(endLiveResp.LiveInfo.Author.StarCount),
		SelfStarCount:  int(endLiveResp.LiveInfo.Author.SelfStarCount),
		SelfLikeCount:  int(endLiveResp.LiveInfo.Author.SelfLikeCount),
		LiveCount:      int(endLiveResp.LiveInfo.Author.LiveCount),
		WorkCount:      int(endLiveResp.LiveInfo.Author.WorkCount),
		FriendCount:    int(endLiveResp.LiveInfo.Author.FriendCount),
		Role:           endLiveResp.LiveInfo.Author.Role,
		Gender:         endLiveResp.LiveInfo.Author.Gender,
		Status:         uint(endLiveResp.LiveInfo.Author.Status),
	}

	return &resp, nil
}

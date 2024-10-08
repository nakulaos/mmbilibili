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

type LiveListByGorseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveListByGorseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveListByGorseLogic {
	return &LiveListByGorseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveListByGorseLogic) LiveListByGorse(req *types.LiveListReq) (resp *types.LiveListResp, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	pagesize := int32(req.PageSize)
	page := int32(req.Page)
	itemid := uint32(req.ItemID)
	total := int32(req.Total)
	in := live.LiveListReq{
		UserId:        uint32(uid),
		PageSize:      &pagesize,
		Page:          &page,
		ItemId:        &itemid,
		Category:      &req.Category,
		Total:         &total,
		RecommendType: &req.RecommendType,
	}
	liveListResp, err := l.svcCtx.LiveBusinessServiceClient.LiveListByGorse(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "live list by gorse error: uid: %d ,pagesize :%d,page:%d,itemid:%d", uid, pagesize, page, itemid)
	}

	resp = &types.LiveListResp{}
	resp.List = make([]types.LiveInfo, 0)
	for i := 0; i < len(liveListResp.List); i++ {
		var author types.User
		author = types.User{
			Id:             liveListResp.List[i].Author.Id,
			Nickname:       liveListResp.List[i].Author.Nickname,
			Username:       liveListResp.List[i].Author.Username,
			Avatar:         liveListResp.List[i].Author.Avatar,
			FollowerCount:  int(liveListResp.List[i].Author.FollowerCount),
			FollowingCount: int(liveListResp.List[i].Author.FollowingCount),
			LikeCount:      int(liveListResp.List[i].Author.LikeCount),
			StarCount:      int(liveListResp.List[i].Author.StarCount),
			SelfStarCount:  int(liveListResp.List[i].Author.SelfStarCount),
			SelfLikeCount:  int(liveListResp.List[i].Author.SelfLikeCount),
			LiveCount:      int(liveListResp.List[i].Author.LiveCount),
			WorkCount:      int(liveListResp.List[i].Author.WorkCount),
			FriendCount:    int(liveListResp.List[i].Author.FriendCount),
			Role:           liveListResp.List[i].Author.Role,
			Gender:         liveListResp.List[i].Author.Gender,
			Status:         uint(liveListResp.List[i].Author.Status),
		}
		var liveInfo types.LiveInfo
		liveInfo = types.LiveInfo{
			LiveID:       uint(liveListResp.List[i].LiveId),
			UserID:       uint(liveListResp.List[i].UserId),
			Title:        liveListResp.List[i].Title,
			Cover:        liveListResp.List[i].Cover,
			Status:       liveListResp.List[i].Status,
			StartTime:    liveListResp.List[i].StartTime,
			EndTime:      liveListResp.List[i].EndTime,
			WatchCount:   int(liveListResp.List[i].WatchCount),
			LikeCount:    int(liveListResp.List[i].LikeCount),
			CommentCount: int(liveListResp.List[i].CommentCount),
			ShareCount:   int(liveListResp.List[i].ShareCount),
			IsLike:       liveListResp.List[i].IsLike,
			IsFollow:     liveListResp.List[i].IsFollow,
			IsStar:       liveListResp.List[i].IsStar,
			IsSelf:       liveListResp.List[i].IsSelf,
			Type:         int(liveListResp.List[i].Type),
			Description:  liveListResp.List[i].Description,
			PlayerUrl:    liveListResp.List[i].PlayerUrl,
			CoverUrl:     liveListResp.List[i].CoverUrl,
			IsOver:       liveListResp.List[i].IsOver,
			Category:     liveListResp.List[i].Category,
			Tags:         liveListResp.List[i].Tags,
			Partition:    liveListResp.List[i].Partition,
			RoomID:       uint(liveListResp.List[i].RoomId),
			Token:        liveListResp.List[i].Token,
		}
		liveInfo.Author = author
		resp.List = append(resp.List, liveInfo)
	}
	resp.Total = int(liveListResp.Total)
	return resp, nil
}

package livebusinessrpcservicelogic

import (
	"backend/common/constant"
	"backend/dao/model"
	"backend/dao/statistics"
	"backend/pkg/xerror"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"github.com/pkg/errors"
	"github.com/zhenghaoz/gorse/client"
	"strconv"
	"strings"

	"backend/rpc/live/internal/svc"
	"backend/rpc/live/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveListByGorseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveListByGorseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveListByGorseLogic {
	return &LiveListByGorseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveListByGorseLogic) LiveListByGorse(in *live.LiveListReq) (*live.LiveListResp, error) {
	uid := in.UserId

	var total, page, pagesize int
	var err error
	if in.Total != nil && (*in.Total > 20 || *in.Total <= 0) {
		total = 20
	}

	var liveIds []uint
	var recommendType, category string
	if in.RecommendType != nil {
		recommendType = *in.RecommendType
	} else {
		recommendType = "popular"
	}

	if in.Category != nil {
		category = *in.Category
	}

	if in.Page == nil {
		page = 0
	} else {
		page = int(*in.Page)
	}

	if in.PageSize == nil {
		pagesize = 10
	} else {
		pagesize = int(*in.PageSize)
	}

	switch recommendType {
	case "recommend":

		if category == "" {
			liveIds, err = RecommendLiveListByGorseWithUid(l.ctx, l.svcCtx.GorseClient, total, page*pagesize, uint(uid))
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] RecommendLiveListByGorseWithUid error: %v", err)
			}
		}
		if category != "" {
			liveIds, err = RecommendLiveListByGorseWithUidAndCategory(l.ctx, l.svcCtx.GorseClient, total, page*pagesize, uint(uid), category)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] RecommendLiveListByGorseWithUidAndCategory error: %v", err)
			}
		}
	case "popular":
		if category == "" {
			liveIds, err = popularLiveListByGorse(l.ctx, l.svcCtx.GorseClient, uint(uid), total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] popularLiveListByGorse error: %v", err)
			}
		}
		if category != "" {
			liveIds, err = popularLiveListByGorseWithCategory(l.ctx, l.svcCtx.GorseClient, uint(uid), category, total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] popularLiveListByGorseWithCategory error: %v", err)
			}
		}
	case "latest":
		if category == "" {
			liveIds, err = latestLiveListByGorse(l.ctx, l.svcCtx.GorseClient, uint(uid), total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] latestLiveListByGorse error: %v", err)
			}
		}
		if category != "" {
			liveIds, err = latestLiveListByGorseWithCategory(l.ctx, l.svcCtx.GorseClient, uint(uid), category, total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] latestLiveListByGorseWithCategory error: %v", err)
			}
		}
	case "neighbors":
		if category == "" {
			if in.ItemId == nil {
				return nil, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"Params": "ItemID"}), "[LiveListByGorse] ItemId is empty")
			}
			liveIds, err = neighborsLiveListByGorse(l.ctx, l.svcCtx.GorseClient, uint(uid), uint(*in.ItemId), total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] neighborsLiveListByGorse error: %v", err)
			}
		}
		if category != "" {
			if in.ItemId == nil {
				return nil, errors.Wrapf(xerror.InvalidParamsError.WithTemplateData(map[string]interface{}{"Params": "ItemID"}), "[LiveListByGorse] ItemId is empty")
			}
			liveIds, err = neighborsLiveListByGorseWithCategory(l.ctx, l.svcCtx.GorseClient, uint(uid), uint(*in.ItemId), category, total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] neighborsLiveListByGorseWithCategory error: %v", err)
			}
		}
	default:
		if category == "" {
			liveIds, err = popularLiveListByGorse(l.ctx, l.svcCtx.GorseClient, uint(uid), total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] popularLiveListByGorse error: %v", err)
			}
		}
		if category != "" {
			liveIds, err = popularLiveListByGorseWithCategory(l.ctx, l.svcCtx.GorseClient, uint(uid), category, total, page*pagesize)
			if err != nil {
				return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] popularLiveListByGorseWithCategory error: %v", err)
			}
		}
	}

	// 获取直播列表
	// 先从缓存获取
	var notfindLiveIds []uint
	var findLiveIds []uint
	var liveModels []*model.Live
	for _, liveId := range liveIds {
		var liveModel model.Live
		err = l.svcCtx.Cache.Get(l.ctx, constant.LiveInfoCacheKey+fmt.Sprintf("%d", liveId), &liveModel)
		if err != nil {
			notfindLiveIds = append(notfindLiveIds, liveId)
			continue
		}

		// 查用户表和直播表，看看有没有like记录
		// 用户是uid，直播是liveModel中的lid
		liveModels = append(liveModels, &liveModel)
		findLiveIds = append(findLiveIds, liveId)

	}

	// 获取是否点赞
	ret, err := l.svcCtx.Dao.User.WithContext(l.ctx).Preload(l.svcCtx.Dao.User.FavoriteLives.On(l.svcCtx.Dao.Live.ID.In(findLiveIds...))).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get user favorite live failed: %v", err)
	}
	var mIsLikeId map[uint]bool = make(map[uint]bool)
	for _, r := range ret.FavoriteLives {
		mIsLikeId[r.ID] = true
	}

	// 从数据库获取,并缓存数据
	if len(notfindLiveIds) > 0 {
		notfindliveModels, err := l.svcCtx.Dao.Live.WithContext(l.ctx).
			Preload(l.svcCtx.Dao.Live.Categories).
			Preload(l.svcCtx.Dao.Live.FavoriteUser.On(l.svcCtx.Dao.User.ID.Eq(uint(uid)))).
			Where(l.svcCtx.Dao.Live.ID.In(notfindLiveIds...)).Find()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get live list failed: %v", err)
		}
		for _, liveModel := range notfindliveModels {
			if liveModel.FavoriteUser != nil {
				mIsLikeId[liveModel.ID] = true
			}
			liveModels = append(liveModels, liveModel)
		}

		go func() {
			for _, liveModel := range notfindliveModels {
				err := l.svcCtx.Cache.Set(&cache.Item{
					Key:   constant.LiveInfoCacheKey + fmt.Sprintf("%d", liveModel.ID),
					Value: liveModel,
				})
				if err != nil {
					l.Logger.Errorf("[LiveListByGorse] set live info cache failed: %v", err)
				}
			}
		}()

	}

	// 预加载一下
	// 查用户表和直播表，看看有没有like记录
	// 用户是uid，直播是liveModel中的lid

	// 获取用户信息
	//先查缓存
	var notfindUserIds []uint
	var findUserIds []uint
	var userIds []uint
	var userModels map[uint]*model.User = make(map[uint]*model.User)
	for _, liveModel := range liveModels {
		var userModel model.User
		err = l.svcCtx.Cache.Get(l.ctx, constant.UserInfoCacheUidKey+fmt.Sprintf("%d", liveModel.UID), &userModel)
		if err != nil {
			notfindUserIds = append(notfindUserIds, liveModel.UID)
			continue
		}
		userModels[liveModel.UID] = &userModel
		findUserIds = append(findUserIds, liveModel.UID)
	}
	userIds = append(userIds, findUserIds...)
	userIds = append(userIds, notfindUserIds...)
	// 从数据库获取,并缓存数据
	if len(notfindUserIds) > 0 {
		/*
			// Followers: 用户被哪些人关注
			// Following: 用户关注了哪些人
			Followers []*User `gorm:"many2many:user_follows;joinForeignKey:FollowedID;joinReferences:FollowerID"`
			// joinForeignKey: 当前用户是被关注的用户，对应的是FollowedID。
			// joinReferences: 关注当前用户的人，其ID是FollowerID。

			Following []*User `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowedID"`
			// joinForeignKey: 当前用户是关注其他用户的人，对应的是FollowerID。
			// joinReferences: 被当前用户关注的人，其ID是FollowedID。
		*/
		notfindUserModels, err := l.svcCtx.Dao.User.WithContext(l.ctx).
			//Preload(l.svcCtx.Dao.User.FavoriteLives.On(l.svcCtx.Dao.Live.ID.Eq())).
			Where(l.svcCtx.Dao.User.ID.In(notfindUserIds...)).Find()
		if err != nil {
			return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get user list failed: %v", err)
		}
		for _, userModel := range notfindUserModels {
			userModels[userModel.ID] = userModel
		}
		go func() {
			for _, userModel := range notfindUserModels {
				err := l.svcCtx.Cache.Set(&cache.Item{
					Key:   constant.UserInfoCacheUidKey + fmt.Sprintf("%d", userModel.ID),
					Value: userModel,
				})
				if err != nil {
					l.Logger.Infof("[LiveListByGorse] set user info cache failed: %v", err)
				}
				err = l.svcCtx.Cache.Set(&cache.Item{
					Key:   constant.UserNameToUidKey + userModel.Username,
					Value: userModel.ID,
				})

				if err != nil {
					l.Logger.Errorf("[LoginWithUsername] set user info cache failed: %v", err)
				}
			}
		}()
	}

	// 获取用户数据
	userStats, err := statistics.GetUsersRelevantCount(l.svcCtx.Redis, l.ctx, userIds)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get user relevant count failed: %v", err)
	}
	// 获取直播数据
	liveStats, err := statistics.GetLivesRelevantCount(l.svcCtx.Redis, l.ctx, liveIds)
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get live relevant count failed: %v", err)
	}

	//
	requestUser, err := l.svcCtx.Dao.User.WithContext(l.ctx).Preload(l.svcCtx.Dao.User.Following.On(l.svcCtx.Dao.User.ID.In(userIds...))).Where(l.svcCtx.Dao.User.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, errors.Wrapf(xerror.ServerError, "[LiveListByGorse] get user info failed: %v", err)
	}
	var mIsFollowId map[uint]bool = make(map[uint]bool)
	for _, r := range requestUser.Following {
		mIsFollowId[r.ID] = true
	}
	// 构建返回数据
	resp := &live.LiveListResp{}
	resp.Total = int32(len(liveModels))
	for _, liveModel := range liveModels {
		var liveInfo live.LiveInfo
		var isself bool
		if uint32(liveModel.UID) == uid {
			isself = true
		}
		liveStat := liveStats[liveModel.ID]
		liveInfo.LiveId = uint32(liveModel.ID)
		liveInfo.UserId = uint32(liveModel.UID)
		liveInfo.Title = liveModel.Title
		liveInfo.Cover = liveModel.CoverURL
		liveInfo.Status = uint32(liveModel.Status)
		liveInfo.StartTime = liveModel.StartTime.Unix()
		liveInfo.Description = liveModel.Description
		liveInfo.WatchCount = int32(liveStat.ViewCount)
		liveInfo.LikeCount = int32(liveStat.LikeCount)
		liveInfo.CommentCount = int32(liveStat.CommentCount)
		liveInfo.ShareCount = 0
		liveInfo.IsLike = mIsLikeId[liveModel.ID]
		liveInfo.IsSelf = isself
		liveInfo.IsFollow = mIsFollowId[liveModel.UID]
		for _, category := range liveModel.Categories {
			liveInfo.Category = append(liveInfo.Category, category.Name)
		}
		liveInfo.PlayerUrl = liveModel.PlayURL
		liveInfo.CoverUrl = liveModel.CoverURL
		liveInfo.Partition = liveModel.Partition
		for _, tag := range liveModel.Tags {
			liveInfo.Tags = append(liveInfo.Tags, tag.Name)
		}
		liveInfo.RoomId = uint32(liveModel.RoomID)
		liveInfo.Token = liveModel.PlayToken
		liveInfo.IsOver = liveModel.IsOver == 1
		liveInfo.Author = &live.User{
			Id:             uint32(userModels[liveModel.UID].ID),
			Username:       userModels[liveModel.UID].Username,
			Nickname:       userModels[liveModel.UID].Nickname,
			Avatar:         userModels[liveModel.UID].Avatar,
			Gender:         uint32(userModels[liveModel.UID].Gender),
			Role:           uint32(userModels[liveModel.UID].Role),
			FollowerCount:  int32(userStats[liveModel.UID].FollowerCount),
			FollowingCount: int32(userStats[liveModel.UID].FollowingCount),
			LikeCount:      int32(userStats[liveModel.UID].LikeCount),
			StarCount:      int32(userStats[liveModel.UID].StarCount),
			LiveCount:      int32(userStats[liveModel.UID].LiveCount),
			WorkCount:      int32(userStats[liveModel.UID].WorkCount),
			FriendCount:    int32(userStats[liveModel.UID].FriendCount),
			Status:         uint32(userModels[liveModel.UID].Status),
		}
		resp.List = append(resp.List, &liveInfo)
	}

	return resp, nil
}

func RecommendLiveListByGorseWithUid(ctx context.Context, gorseClient *client.GorseClient, total, offset int, uid uint) (resp []uint, err error) {

	liveGorseItem, err := gorseClient.GetItemRecommend(ctx, fmt.Sprintf("%d", uid), []string{}, "", "", total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func RecommendLiveListByGorseWithUidAndCategory(ctx context.Context, gorseClient *client.GorseClient, total, offset int, uid uint, category string) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemRecommendWithCategory(ctx, fmt.Sprintf("%d", uid), category, "", "", total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func popularLiveListByGorse(ctx context.Context, gorseClient *client.GorseClient, uid uint, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemPopular(ctx, fmt.Sprintf("%d", uid), total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func popularLiveListByGorseWithCategory(ctx context.Context, gorseClient *client.GorseClient, uid uint, category string, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemPopularWithCategory(ctx, fmt.Sprintf("%d", uid), category, total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func latestLiveListByGorse(ctx context.Context, gorseClient *client.GorseClient, uid uint, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemLatest(ctx, fmt.Sprintf("%d", uid), total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func latestLiveListByGorseWithCategory(ctx context.Context, gorseClient *client.GorseClient, uid uint, category string, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemLatestWithCategory(ctx, fmt.Sprintf("%d", uid), category, total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func neighborsLiveListByGorse(ctx context.Context, gorseClient *client.GorseClient, uid, itemId uint, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemNeighbors(ctx, fmt.Sprintf("%s:%d", constant.LiveType, itemId), fmt.Sprintf("%d", uid), total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

func neighborsLiveListByGorseWithCategory(ctx context.Context, gorseClient *client.GorseClient, uid, itemId uint, category string, total, offset int) (resp []uint, err error) {
	liveGorseItem, err := gorseClient.GetItemNeighborsWithCategory(ctx, fmt.Sprintf("%s:%d", constant.LiveType, itemId), category, fmt.Sprintf("%d", uid), total, offset)
	if err != nil {
		return resp, err
	}
	// 分离
	var liveIds []uint
	for _, gorseItem := range liveGorseItem {
		item := strings.Split(gorseItem.Id, ":")
		liveId, _ := strconv.Atoi(item[1])
		liveIds = append(liveIds, uint(liveId))
	}
	return liveIds, nil
}

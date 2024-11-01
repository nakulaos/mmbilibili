package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/global"
	"backend/app/rpc/user/biz/model"
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/segmentio/kafka-go"
	"time"
)

type AddFollowingService struct {
	ctx context.Context
} // NewAddFollowingService new AddFollowingService
func NewAddFollowingService(ctx context.Context) *AddFollowingService {
	return &AddFollowingService{ctx: ctx}
}

// Run create note info
func (s *AddFollowingService) Run(req *user.AddFollowingReq) (resp *user.AddFollowingResp, err error) {
	/*
		1. 判断uid和fid是否为空
		2. 判断uid和fid是否相等
		3. 获取关注信息

	*/
	var (
		uid              = req.Uid
		rid              = req.Rid
		ur               *model.UserRelationship
		rr               *model.UserRelationship
		ust              *model.UserRelevantCount
		umsg             = model.NewUserRelevantCountMessage(uid)
		rmsg             = model.NewUserRelevantCountMessage(rid)
		maxFollowerCount = global.Config.App.FollowingMaxCount
		friend           bool
	)

	if uid <= 0 || rid <= 0 {
		return nil, ecode.InvalidParamsError.WithTemplateData(map[string]string{"Params": "uid or fid"})
	}

	if uid == rid {
		klog.Errorf("addFollowingService Run uid == fid")
		return nil, ecode.InvalidParamsError.WithTemplateData(map[string]string{"Params": "uid or fid"})
	}

	if ur, err = global.UserDal.GetOrCreateMidFidRelation(s.ctx, uid, rid); err != nil {
		return nil, ecode.ServerError
	}

	if rr, err = global.UserDal.GetOrCreateMidFidRelation(s.ctx, rid, uid); err != nil {
		return nil, ecode.ServerError
	}

	marrt := model.RelationshipAttr(ur.RelationshipAttr)
	rarrt := model.RelationshipAttr(rr.RelationshipAttr)
	if ust, err = global.UserDal.GetUserRelevantCountByID(s.ctx, uid); err != nil {
		return nil, ecode.ServerError
	}

	switch rarrt {
	case model.RelationshipAttrFollowing: // 已经关注
		// 朋友关系
		ur.RelationshipAttr = model.SetAttr(ur.RelationshipAttr, model.RelationshipAttrFriend)
		rr.RelationshipAttr = model.SetAttr(rr.RelationshipAttr, model.RelationshipAttrFriend)
		friend = true
	}

	switch marrt {
	case model.RelationshipAttrBlack:
		// 黑名单
		err = ecode.UserBlackUserError
		return nil, err
	case model.RelationshipAttrFriend:
		err = ecode.UserAlreadyFriendError
		return nil, err
	case model.RelationshipAttrFollowing:
		err = ecode.UserAlreadyFollowingError
		return nil, err
	case model.RelationshipAttrWhisper:
		// 直接关注,关系升级
		if ust.FollowingCount > maxFollowerCount {
			err = ecode.UserFollowingMaxError
			return nil, err
		}

		if friend {
			// 更新两个关系
			if err = global.UserDal.UpdateFriendRelation(s.ctx, uid, rid, maxFollowerCount); err != nil {
				return nil, ecode.ServerError
			}
		} else {
			if err = global.UserDal.UpdateMidRelation(s.ctx, uid, rid, model.RelationshipAttrFollowing, maxFollowerCount); err != nil {
				return nil, err
			}
		}

		umsg.CountChange[model.TypeWhisperCount] = -1
		umsg.CountChange[model.TypeFollowingCount] = 1
		rmsg.CountChange[model.TypeFollowerCount] = 1
		if friend {
			umsg.CountChange[model.TypeFriendCount] = 1
			rmsg.CountChange[model.TypeFriendCount] = 1
		}

	case model.RelationshipAttrNone:
		// 直接关注
		if ust.FollowingCount > maxFollowerCount {
			err = ecode.UserFollowingMaxError
			return nil, err
		}

		if friend {
			// 更新两个关系
			if err = global.UserDal.UpdateFriendRelation(s.ctx, uid, rid, maxFollowerCount); err != nil {
				return nil, ecode.ServerError
			}
		} else {
			if err = global.UserDal.UpdateMidRelation(s.ctx, uid, rid, model.RelationshipAttrFollowing, maxFollowerCount); err != nil {
				return nil, err
			}
		}

		if err = global.UserDal.UpdateMidRelation(s.ctx, uid, rid, model.RelationshipAttrFollowing, maxFollowerCount); err != nil {
			return nil, err
		}

		umsg.CountChange[model.TypeFollowingCount] = 1
		rmsg.CountChange[model.TypeFollowerCount] = 1
		if friend {
			umsg.CountChange[model.TypeFriendCount] = 1
			rmsg.CountChange[model.TypeFriendCount] = 1
		}

	}

	go func() {
		for i := 1; i <= 10; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err = global.UserRelevantCountProducer.WriteMessages(ctx, kafka.Message{
				Key:   umsg.GetUserRelevantCountMessageKey(),
				Value: umsg.Json(),
			}); err != nil {
				klog.Errorf("global.UserRelevantCountProducer.WriteMessages(%d) err:%v", umsg.GetUserRelevantCountMessageKey(), err)
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}
		for i := 1; i <= 10; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err = global.UserRelevantCountProducer.WriteMessages(ctx, kafka.Message{
				Key:   rmsg.GetUserRelevantCountMessageKey(),
				Value: rmsg.Json(),
			}); err != nil {
				klog.Errorf("global.UserRelevantCountProducer.WriteMessages(%d) err:%v", umsg.GetUserRelevantCountMessageKey(), err)
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}

		// todo 添加关注通知

	}()

	return
}

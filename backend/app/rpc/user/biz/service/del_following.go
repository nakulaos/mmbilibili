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

type DelFollowingService struct {
	ctx context.Context
} // NewDelFollowingService new DelFollowingService
func NewDelFollowingService(ctx context.Context) *DelFollowingService {
	return &DelFollowingService{ctx: ctx}
}

// Run create note info
func (s *DelFollowingService) Run(req *user.DelFollowingReq) (resp *user.DelFollowingResp, err error) {
	var (
		uid = req.Uid
		rid = req.Rid
		ur  *model.UserRelationship
		//rr   *model.UserRelationship
		umsg             = model.NewUserRelevantCountMessage(uid)
		rmsg             = model.NewUserRelevantCountMessage(rid)
		maxFollowerCount = global.Config.App.FollowingMaxCount

		friend bool
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

	if _, err = global.UserDal.GetOrCreateMidFidRelation(s.ctx, rid, uid); err != nil {
		return nil, ecode.ServerError
	}

	marrt := model.RelationshipAttr(ur.RelationshipAttr)
	//rarrt := model.RelationshipAttr(rr.RelationshipAttr)
	switch marrt {
	case model.RelationshipAttrFriend:
		friend = true
	case model.RelationshipAttrNone:
		return nil, ecode.UserAlreadyNotFollowingError
	case model.RelationshipAttrFollowing:
	default:
		// TODO 缓存删除
	}

	if friend {
		if err = global.UserDal.RemoveFriendRelation(s.ctx, uid, rid, maxFollowerCount); err != nil {
			return nil, ecode.ServerError
		}
		umsg.CountChange[model.TypeFollowingCount] = -1
		umsg.CountChange[model.TypeFriendCount] = -1
		rmsg.CountChange[model.TypeFollowerCount] = -1
		rmsg.CountChange[model.TypeFriendCount] = -1
	} else {
		if err = global.UserDal.UpdateMidRelation(s.ctx, uid, rid, model.RelationshipAttrNone, maxFollowerCount); err != nil {
			return nil, ecode.ServerError
		}
		umsg.CountChange[model.TypeFollowingCount] = -1
		rmsg.CountChange[model.TypeFollowerCount] = -1
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
				time.Sleep(1 * time.Second) // 可选：添加重试间隔
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
	}()

	return &user.DelFollowingResp{}, nil
}

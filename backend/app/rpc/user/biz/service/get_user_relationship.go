package service

import (
	"backend/app/rpc/user/biz/global"
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	"time"
)

type GetUserRelationshipService struct {
	ctx context.Context
} // NewGetUserRelationshipService new GetUserRelationshipService
func NewGetUserRelationshipService(ctx context.Context) *GetUserRelationshipService {
	return &GetUserRelationshipService{ctx: ctx}
}

// Run create note info
func (s *GetUserRelationshipService) Run(req *user.GetUserRelationshipReq) (resp *user.GetUserRelationshipResp, err error) {
	var (
		uid               = req.Uid
		rids              = req.Rids
		followingExpire   = time.Duration(global.Config.App.FollowingExpire) * time.Second
		followingMaxCount = global.Config.App.FollowingMaxCount
	)

	if uid <= 0 || len(rids) == 0 {
		return
	}

	// 获取用户关系
	resp = &user.GetUserRelationshipResp{}
	resp.List = make([]*user.UserRelationship, 0)

	relationships, err := global.UserDal.GetAttrsByUIDAndRIDS(s.ctx, uid, rids, followingMaxCount, followingExpire)
	for _, rid := range rids {
		if _, ok := relationships[rid]; !ok {
			resp.List = append(resp.List, &user.UserRelationship{
				Rid:  rid,
				Attr: relationships[rid].RelationshipAttr,
			})
		}
	}

	return resp, err
}

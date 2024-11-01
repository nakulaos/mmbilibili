package service

import (
	"backend/app/common/ecode"
	"backend/app/rpc/user/biz/global"
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

type FollowerListService struct {
	ctx context.Context
} // NewFollowerListService new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

// Run create note info
func (s *FollowerListService) Run(req *user.FollowerListReq) (resp *user.FollowerListResp, err error) {
	var (
		limit  = req.Total
		offset = (req.Page - 1) * req.PageSize
	)
	if limit <= 0 || limit > _maxLimit {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	users, err := global.UserDal.GetFollowersByUserID(s.ctx, req.ActionId, limit, offset)
	if err != nil {
		return nil, ecode.ServerError
	}

	resp = &user.FollowerListResp{}
	resp.List = make([]*user.User, 0, len(users))
	for _, u := range users {
		resp.List = append(resp.List, &user.User{
			Avatar:   u.Avatar,
			Id:       u.ID,
			Nickname: u.Nickname,
			Username: u.Username,
			Gender:   uint32(u.Gender),
			Role:     uint32(u.Role),
			Status:   uint32(u.Status),
		})
	}
	return
}

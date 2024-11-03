package service

import (
	"backend/app/http/biz/global"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FriendListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFriendListService(Context context.Context, RequestContext *app.RequestContext) *FriendListService {
	return &FriendListService{RequestContext: RequestContext, Context: Context}
}

func (h *FriendListService) Run(req *user.FriendListReq) (resp *user.FriendListResp, err error) {
	var (
		uid      = req.ActionID
		page     = req.Page
		pageSize = req.PageSize
		total    = req.Total
	)

	friendListResp, err := global.UserRpcClient.FriendList(h.Context, &userRpc.FriendListReq{
		ActionId: uid,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})

	if err != nil {
		return nil, err
	}

	resp = &user.FriendListResp{}
	resp.Total = friendListResp.Total
	resp.List = make([]*user.User, 0)
	for _, v := range friendListResp.List {
		userInfo := &user.User{
			Id:       v.Id,
			Username: v.Username,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Status:   int64(v.Status),
			Role:     int64(v.Role),
			Gender:   int64(v.Gender),
		}
		resp.List = append(resp.List, userInfo)
	}

	return resp, nil
}

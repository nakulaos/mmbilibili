package service

import (
	"backend/app/http/biz/global"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowingListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowingListService(Context context.Context, RequestContext *app.RequestContext) *FollowingListService {
	return &FollowingListService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowingListService) Run(req *user.FollowingListReq) (resp *user.FollowingListResp, err error) {
	var (
		uid      = tools.GetUserID(h.RequestContext)
		page     = req.Page
		pageSize = req.PageSize
		total    = req.Total
	)

	followingListResp, err := global.UserRpcClient.FollowingList(h.Context, &userRpc.FollowingListReq{
		ActionId: uid,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})

	if err != nil {
		return nil, err
	}

	resp = &user.FollowingListResp{}
	resp.Total = followingListResp.Total
	resp.List = make([]*user.User, 0)
	for _, v := range followingListResp.List {
		userInfo := &user.User{
			Id:       v.Id,
			Username: v.Username,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Status:   int64(v.Status),
			Gender:   int64(v.Gender),
			Role:     int64(v.Role),
		}
		resp.List = append(resp.List, userInfo)
	}

	return resp, nil
}

package service

import (
	"backend/app/http/biz/global"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowerListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowerListService(Context context.Context, RequestContext *app.RequestContext) *FollowerListService {
	return &FollowerListService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowerListService) Run(req *user.FollowerListReq) (resp *user.FollowerListResp, err error) {
	var (
		uid      = tools.GetUserID(h.RequestContext)
		page     = req.Page
		pageSize = req.PageSize
		total    = req.Total
	)

	followerListResp, err := global.UserRpcClient.FollowerList(h.Context, &userRpc.FollowerListReq{
		ActionId: uid,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})

	if err != nil {
		return nil, err
	}

	resp = &user.FollowerListResp{}
	resp.Total = followerListResp.Total
	resp.List = make([]*user.User, 0)
	for _, v := range followerListResp.List {
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

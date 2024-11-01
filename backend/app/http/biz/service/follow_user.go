package service

import (
	"backend/app/common/constant"
	"backend/app/http/biz/global"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"context"

	user "backend/app/http/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FollowUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowUserService(Context context.Context, RequestContext *app.RequestContext) *FollowUserService {
	return &FollowUserService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowUserService) Run(req *user.FollowUserReq) (resp *user.FollowUserResp, err error) {
	uid := tools.GetUserID(h.RequestContext)
	rid := req.UserID
	action := req.Action

	if action == constant.FollowAction {
		_, err = global.UserRpcClient.AddFollowing(h.Context, &userRpc.AddFollowingReq{
			Uid: uid,
			Rid: rid,
		})
	} else {
		_, err = global.UserRpcClient.DelFollowing(h.Context, &userRpc.DelFollowingReq{
			Uid: uid,
			Rid: rid,
		})
	}

	return &user.FollowUserResp{}, err
}

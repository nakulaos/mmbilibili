package service

import (
	"backend/app/http/biz/global"
	user "backend/app/http/hertz_gen/user"
	userRpc "backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	updateUserInfoReq := buildUpdateUserInfoReq(req)

	updateUserInfoReq.Id = uint32(tools.GetUserID(h.RequestContext))

	updateUserInfoResp, err := global.UserRpcClient.UpdateUserInfo(h.Context, updateUserInfoReq)
	if err != nil {
		hlog.Errorf("UpdateUserInfoService.GlobalUserRpcClient.UpdateUserInfo(req:%v) error: %v", updateUserInfoReq, err)
		return nil, err
	}

	resp = &user.UpdateUserInfoResp{}
	resp.UserInfo = &user.User{}
	resp.UserInfo.Id = updateUserInfoResp.UserInfo.Id
	resp.UserInfo.Username = updateUserInfoResp.UserInfo.Username
	resp.UserInfo.Nickname = updateUserInfoResp.UserInfo.Nickname
	resp.UserInfo.Avatar = updateUserInfoResp.UserInfo.Avatar
	resp.UserInfo.Status = int64(updateUserInfoResp.UserInfo.Status)
	resp.UserInfo.Gender = int64(updateUserInfoResp.UserInfo.Gender)
	resp.UserInfo.Role = int64(updateUserInfoResp.UserInfo.Role)
	resp.UserInfo.FollowerCount = updateUserInfoResp.UserInfo.FollowerCount
	resp.UserInfo.FollowingCount = updateUserInfoResp.UserInfo.FollowingCount
	resp.UserInfo.LikeCount = updateUserInfoResp.UserInfo.LikeCount
	resp.UserInfo.StarCount = updateUserInfoResp.UserInfo.StarCount
	resp.UserInfo.SelfStarCount = updateUserInfoResp.UserInfo.SelfStarCount
	resp.UserInfo.SelfLikeCount = updateUserInfoResp.UserInfo.SelfLikeCount
	resp.UserInfo.LiveCount = updateUserInfoResp.UserInfo.LiveCount
	resp.UserInfo.WorkCount = updateUserInfoResp.UserInfo.WorkCount
	resp.UserInfo.FriendCount = updateUserInfoResp.UserInfo.FriendCount
	resp.UserInfo.Phone = updateUserInfoResp.UserInfo.Phone
	resp.UserInfo.Email = updateUserInfoResp.UserInfo.Email

	return
}

func buildUpdateUserInfoReq(req *user.UpdateUserInfoReq) *userRpc.UpdateUserInfoReq {
	rreq := &userRpc.UpdateUserInfoReq{}
	if req.Nickname != "" {
		rreq.Nickname = &req.Nickname
	}
	if req.Avatar != "" {
		rreq.Avatar = &req.Avatar
	}
	if req.Phone != "" {
		rreq.Phone = &req.Phone
	}
	if req.Email != "" {
		rreq.Email = &req.Email
	}
	if req.Gender != 0 {
		gender := uint32(req.Gender)
		rreq.Gender = &gender
	}
	if req.Role != 0 {
		role := uint32(req.Role)
		rreq.Role = &role
	}
	return rreq
}

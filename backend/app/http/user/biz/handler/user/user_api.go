package user

import (
	"context"

	"backend/app/http/user/biz/service"
	"backend/app/http/user/biz/utils"
	user "backend/app/http/user/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LoginWithUsername .
// @router /login/username [POST]
func LoginWithUsername(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginWithUsernameReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.LoginResp{}
	resp, err = service.NewLoginWithUsernameService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// LoginWithEmail .
// @router /login/email [POST]
func LoginWithEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginWithEmailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.LoginResp{}
	resp, err = service.NewLoginWithEmailService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// LoginWithPhone .
// @router /login/phone [POST]
func LoginWithPhone(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginWithPhoneReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.LoginResp{}
	resp, err = service.NewLoginWithPhoneService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.LoginResp{}
	resp, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateUserInfo .
// @router /user/info [PUT]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.UpdateUserInfoResp{}
	resp, err = service.NewUpdateUserInfoService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Logout .
// @router /logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LogoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.LogoutResp{}
	resp, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// FollowUser .
// @router /user/follow [POST]
func FollowUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.FollowUserResp{}
	resp, err = service.NewFollowUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// FollowerList .
// @router /user/followers [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowerListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.FollowerListResp{}
	resp, err = service.NewFollowerListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// FollowingList .
// @router /user/following [GET]
func FollowingList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowingListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.FollowingListResp{}
	resp, err = service.NewFollowingListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// FriendList .
// @router /user/friends [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.FriendListResp{}
	resp, err = service.NewFriendListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserUploadFile .
// @router /user/upload [POST]
func UserUploadFile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserUploadFileReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.UserUploadFileResp{}
	resp, err = service.NewUserUploadFileService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

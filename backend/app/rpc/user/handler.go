package main

import (
	"backend/app/rpc/user/biz/service"
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
)

// UserRpcServiceImpl implements the last service interface defined in the IDL.
type UserRpcServiceImpl struct{}

// LoginWithUsername implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) LoginWithUsername(ctx context.Context, req *user.LoginWithUsernameReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginWithUsernameService(ctx).Run(req)

	return resp, err
}

// LoginWithEmail implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) LoginWithEmail(ctx context.Context, req *user.LoginWithEmailReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginWithEmailService(ctx).Run(req)

	return resp, err
}

// LoginWithPhone implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) LoginWithPhone(ctx context.Context, req *user.LoginWithPhoneReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginWithPhoneService(ctx).Run(req)

	return resp, err
}

// Register implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// UpdateUserInfo implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = service.NewUpdateUserInfoService(ctx).Run(req)

	return resp, err
}

// Logout implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) Logout(ctx context.Context, req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}

// FollowUser implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) FollowUser(ctx context.Context, req *user.FollowUserReq) (resp *user.FollowUserResp, err error) {
	resp, err = service.NewFollowUserService(ctx).Run(req)

	return resp, err
}

// FollowerList implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) FollowerList(ctx context.Context, req *user.FollowerListReq) (resp *user.FollowerListResp, err error) {
	resp, err = service.NewFollowerListService(ctx).Run(req)

	return resp, err
}

// FollowingList implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) FollowingList(ctx context.Context, req *user.FollowingListReq) (resp *user.FollowingListResp, err error) {
	resp, err = service.NewFollowingListService(ctx).Run(req)

	return resp, err
}

// FriendList implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) FriendList(ctx context.Context, req *user.FriendListReq) (resp *user.FriendListResp, err error) {
	resp, err = service.NewFriendListService(ctx).Run(req)

	return resp, err
}

// UserUploadFile implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) UserUploadFile(ctx context.Context, req *user.UserUploadFileReq) (resp *user.UserUploadFileResp, err error) {
	resp, err = service.NewUserUploadFileService(ctx).Run(req)

	return resp, err
}

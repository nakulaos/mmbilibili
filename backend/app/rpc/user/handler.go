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

// RefreshToken implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) RefreshToken(ctx context.Context, req *user.RefreshTokenReq) (resp *user.RefreshTokenResp, err error) {
	resp, err = service.NewRefreshTokenService(ctx).Run(req)

	return resp, err
}

// AddFollowing implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) AddFollowing(ctx context.Context, req *user.AddFollowingReq) (resp *user.AddFollowingResp, err error) {
	resp, err = service.NewAddFollowingService(ctx).Run(req)

	return resp, err
}

// AddWhisper implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) AddWhisper(ctx context.Context, req *user.AddWhisperReq) (resp *user.AddWhisperResp, err error) {
	resp, err = service.NewAddWhisperService(ctx).Run(req)

	return resp, err
}

// AddBlack implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) AddBlack(ctx context.Context, req *user.AddBlackReq) (resp *user.AddBlackResp, err error) {
	resp, err = service.NewAddBlackService(ctx).Run(req)

	return resp, err
}

// DelFollowing implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) DelFollowing(ctx context.Context, req *user.DelFollowingReq) (resp *user.DelFollowingResp, err error) {
	resp, err = service.NewDelFollowingService(ctx).Run(req)

	return resp, err
}

// DelWhisper implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) DelWhisper(ctx context.Context, req *user.DelWhisperReq) (resp *user.DelWhisperResp, err error) {
	resp, err = service.NewDelWhisperService(ctx).Run(req)

	return resp, err
}

// DelBlack implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) DelBlack(ctx context.Context, req *user.DelBlackReq) (resp *user.DelBlackResp, err error) {
	resp, err = service.NewDelBlackService(ctx).Run(req)

	return resp, err
}

// GetUserRelationship implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) GetUserRelationship(ctx context.Context, req *user.GetUserRelationshipReq) (resp *user.GetUserRelationshipResp, err error) {
	resp, err = service.NewGetUserRelationshipService(ctx).Run(req)

	return resp, err
}

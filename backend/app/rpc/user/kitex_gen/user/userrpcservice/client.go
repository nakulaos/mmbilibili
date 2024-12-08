// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userrpcservice

import (
	user "backend/app/rpc/user/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LoginWithUsername(ctx context.Context, Req *user.LoginWithUsernameReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	LoginWithEmail(ctx context.Context, Req *user.LoginWithEmailReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	LoginWithPhone(ctx context.Context, Req *user.LoginWithPhoneReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.UpdateUserInfoResp, err error)
	Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error)
	FollowerList(ctx context.Context, Req *user.FollowerListReq, callOptions ...callopt.Option) (r *user.FollowerListResp, err error)
	FollowingList(ctx context.Context, Req *user.FollowingListReq, callOptions ...callopt.Option) (r *user.FollowingListResp, err error)
	FriendList(ctx context.Context, Req *user.FriendListReq, callOptions ...callopt.Option) (r *user.FriendListResp, err error)
	RefreshToken(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.RefreshTokenResp, err error)
	AddFollowing(ctx context.Context, Req *user.AddFollowingReq, callOptions ...callopt.Option) (r *user.AddFollowingResp, err error)
	AddWhisper(ctx context.Context, Req *user.AddWhisperReq, callOptions ...callopt.Option) (r *user.AddWhisperResp, err error)
	AddBlack(ctx context.Context, Req *user.AddBlackReq, callOptions ...callopt.Option) (r *user.AddBlackResp, err error)
	DelFollowing(ctx context.Context, Req *user.DelFollowingReq, callOptions ...callopt.Option) (r *user.DelFollowingResp, err error)
	DelWhisper(ctx context.Context, Req *user.DelWhisperReq, callOptions ...callopt.Option) (r *user.DelWhisperResp, err error)
	DelBlack(ctx context.Context, Req *user.DelBlackReq, callOptions ...callopt.Option) (r *user.DelBlackResp, err error)
	GetUserRelationship(ctx context.Context, Req *user.GetUserRelationshipReq, callOptions ...callopt.Option) (r *user.GetUserRelationshipResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserRpcServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserRpcServiceClient struct {
	*kClient
}

func (p *kUserRpcServiceClient) LoginWithUsername(ctx context.Context, Req *user.LoginWithUsernameReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginWithUsername(ctx, Req)
}

func (p *kUserRpcServiceClient) LoginWithEmail(ctx context.Context, Req *user.LoginWithEmailReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginWithEmail(ctx, Req)
}

func (p *kUserRpcServiceClient) LoginWithPhone(ctx context.Context, Req *user.LoginWithPhoneReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginWithPhone(ctx, Req)
}

func (p *kUserRpcServiceClient) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserRpcServiceClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.UpdateUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kUserRpcServiceClient) Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Logout(ctx, Req)
}

func (p *kUserRpcServiceClient) FollowerList(ctx context.Context, Req *user.FollowerListReq, callOptions ...callopt.Option) (r *user.FollowerListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, Req)
}

func (p *kUserRpcServiceClient) FollowingList(ctx context.Context, Req *user.FollowingListReq, callOptions ...callopt.Option) (r *user.FollowingListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowingList(ctx, Req)
}

func (p *kUserRpcServiceClient) FriendList(ctx context.Context, Req *user.FriendListReq, callOptions ...callopt.Option) (r *user.FriendListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, Req)
}

func (p *kUserRpcServiceClient) RefreshToken(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.RefreshTokenResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RefreshToken(ctx, Req)
}

func (p *kUserRpcServiceClient) AddFollowing(ctx context.Context, Req *user.AddFollowingReq, callOptions ...callopt.Option) (r *user.AddFollowingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFollowing(ctx, Req)
}

func (p *kUserRpcServiceClient) AddWhisper(ctx context.Context, Req *user.AddWhisperReq, callOptions ...callopt.Option) (r *user.AddWhisperResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddWhisper(ctx, Req)
}

func (p *kUserRpcServiceClient) AddBlack(ctx context.Context, Req *user.AddBlackReq, callOptions ...callopt.Option) (r *user.AddBlackResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddBlack(ctx, Req)
}

func (p *kUserRpcServiceClient) DelFollowing(ctx context.Context, Req *user.DelFollowingReq, callOptions ...callopt.Option) (r *user.DelFollowingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelFollowing(ctx, Req)
}

func (p *kUserRpcServiceClient) DelWhisper(ctx context.Context, Req *user.DelWhisperReq, callOptions ...callopt.Option) (r *user.DelWhisperResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelWhisper(ctx, Req)
}

func (p *kUserRpcServiceClient) DelBlack(ctx context.Context, Req *user.DelBlackReq, callOptions ...callopt.Option) (r *user.DelBlackResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelBlack(ctx, Req)
}

func (p *kUserRpcServiceClient) GetUserRelationship(ctx context.Context, Req *user.GetUserRelationshipReq, callOptions ...callopt.Option) (r *user.GetUserRelationshipResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserRelationship(ctx, Req)
}

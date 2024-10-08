package follow

import (
	"backend/rpc/user/user"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	in := user.UpdateUserInfoReq{
		Id: uint32(uid),
	}

	if req.Gender != 0 {
		in.Gender = &req.Gender
	}
	if req.Nickname != "" {
		in.Nickname = &req.Nickname
	}
	if req.Avatar != "" {
		in.Avatar = &req.Avatar
	}
	if req.Phone != "" {
		in.Phone = &req.Phone
	}
	if req.Email != "" {
		in.Email = &req.Email
	}
	updateUserInfoResp, err := l.svcCtx.UserCommonServiceClient.UpdateUserInfo(l.ctx, &in)
	if err != nil {
		return nil, errors.Wrapf(err, "[UpdateUserInfo] call rpc user.UpdateUserInfo : req:%v", req)
	}

	resp = &types.UpdateUserInfoResp{}
	resp.UserInfo.Id = updateUserInfoResp.UserInfo.Id
	resp.UserInfo.Username = updateUserInfoResp.UserInfo.Username
	resp.UserInfo.Nickname = updateUserInfoResp.UserInfo.Nickname
	resp.UserInfo.Avatar = updateUserInfoResp.UserInfo.Avatar
	resp.UserInfo.Status = uint(updateUserInfoResp.UserInfo.Status)
	resp.UserInfo.Gender = uint32(updateUserInfoResp.UserInfo.Gender)
	resp.UserInfo.Role = uint32(updateUserInfoResp.UserInfo.Role)
	resp.UserInfo.FollowerCount = int(updateUserInfoResp.UserInfo.FollowerCount)
	resp.UserInfo.FollowingCount = int(updateUserInfoResp.UserInfo.FollowingCount)
	resp.UserInfo.LikeCount = int(updateUserInfoResp.UserInfo.LikeCount)
	resp.UserInfo.StarCount = int(updateUserInfoResp.UserInfo.StarCount)
	resp.UserInfo.SelfStarCount = int(updateUserInfoResp.UserInfo.SelfStarCount)
	resp.UserInfo.SelfLikeCount = int(updateUserInfoResp.UserInfo.SelfLikeCount)
	resp.UserInfo.LiveCount = int(updateUserInfoResp.UserInfo.LiveCount)
	resp.UserInfo.WorkCount = int(updateUserInfoResp.UserInfo.WorkCount)
	resp.UserInfo.FriendCount = int(updateUserInfoResp.UserInfo.FriendCount)
	resp.UserInfo.Phone = updateUserInfoResp.UserInfo.Phone
	resp.UserInfo.Email = updateUserInfoResp.UserInfo.Email
	return resp, nil
}

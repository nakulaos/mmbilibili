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

type FollowUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowUserLogic) FollowUser(req *types.FollowUserReq) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	actionId := req.UserID
	action := req.Action
	in := user.FollowUserReq{
		UserId:   uint32(uid),
		ActionId: uint32(actionId),
		Action:   int32(action),
	}

	_, err := l.svcCtx.UserFollowServiceClient.FollowUser(l.ctx, &in)
	if err != nil {
		return errors.Wrapf(err, "[FollowUser] call rpc user.FollowUser : userId:%d, actionId:%d, action:%d", uid, actionId, action)
	}

	return nil
}

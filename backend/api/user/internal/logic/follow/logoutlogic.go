package follow

import (
	"backend/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"net/http"

	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq, r *http.Request) error {
	in := user.LogoutReq{
		LoginToken: r.Header.Get("Authorization"),
	}
	_, err := l.svcCtx.UserCommonServiceClient.Logout(l.ctx, &in)
	if err != nil {
		return errors.Wrapf(err, "[Logout] call rpc user.Logout : req:%v", req)
	}
	return nil
}

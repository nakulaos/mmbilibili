package follow

import (
	"backend/pkg/base"
	"backend/pkg/xvalidator"
	"net/http"

	"backend/api/user/internal/logic/follow"
	"backend/api/user/internal/svc"
	"backend/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		e := xvalidator.Validate.StructCtx(r.Context(), &req)
		if e != nil {
			base.HttpResult(r, w, nil, e)
			return
		}

		l := follow.NewFollowUserLogic(r.Context(), svcCtx)
		err := l.FollowUser(&req)
		if err != nil {
			base.HttpResult(r, w, nil, err)
		} else {
			base.HttpResult(r, w, nil, nil)
		}
	}
}

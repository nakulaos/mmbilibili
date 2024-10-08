package callback

import (
	"backend/pkg/base"
	"backend/pkg/xvalidator"
	"net/http"

	"backend/api/live/internal/logic/callback"
	"backend/api/live/internal/svc"
	"backend/api/live/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OnPublicStreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnPublicStreamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		e := xvalidator.Validate.StructCtx(r.Context(), &req)
		if e != nil {
			base.HttpResult(r, w, nil, e)
			return
		}

		l := callback.NewOnPublicStreamLogic(r.Context(), svcCtx)
		resp, err := l.OnPublicStream(&req)
		if err != nil {
			base.HttpResult(r, w, nil, err)
		} else {
			base.HttpResult(r, w, resp, nil)
		}
	}
}

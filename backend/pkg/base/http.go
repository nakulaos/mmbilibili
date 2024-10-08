package base

import (
	"backend/pkg/xerror"
	"backend/pkg/xvalidator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	xerr "github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
	} else {
		causeErr := errors.Cause(err)
		var errCode int
		var errMsg string
		switch e := causeErr.(type) {
		case xerror.XError:
			errCode = e.Code
			errMsg = e.Error()
		case validator.ValidationErrors:
			errCode = 400
			errMsg = xvalidator.Translate(err)
		default:
			if gstatus, ok := status.FromError(causeErr); ok {
				errCode = int(gstatus.Code())
				errMsg = gstatus.Message()
			}

		}
		logx.WithContext(r.Context()).Errorf("API-ERR:%+v", err)
		xhttp.JsonBaseResponseCtx(r.Context(), w, xerr.New(int(errCode), errMsg))
	}
}

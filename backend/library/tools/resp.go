package tools

import (
	"backend/library/metric"
	"backend/library/xerror"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	hertzI18n "github.com/hertz-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	c.JSON(code, Response{
		Code: code,
		Msg:  "ok",
		Data: data,
	})
}

func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	causeErr := err
	var (
		errcode int
		errmsg  string
	)

	switch e := causeErr.(type) {
	case *xerror.XError:
		if e.Code == 100003 {
			code = http.StatusBadRequest
		}
		// 埋点上报
		if e.Code >= 200000 {
			metric.IncrGauge(metric.BusinessError, ConvertToLowerSnakeCase(e.Message.ID))
		}

		errcode = int(e.BizStatusCode())
		errmsg = hertzI18n.MustGetMessage(ctx, &i18n.LocalizeConfig{
			MessageID:    e.BizMessage(),
			TemplateData: e.BizExtra(),
		})
	case *kerrors.BizStatusError:
		if e.BizStatusCode() == 100003 {
			code = http.StatusBadRequest
		}
		// 埋点上报
		if e.BizStatusCode() >= 200000 {
			metric.IncrGauge(metric.BusinessError, ConvertToLowerSnakeCase(e.BizMessage()))
		}

		errcode = int(e.BizStatusCode())
		errmsg = hertzI18n.MustGetMessage(ctx, &i18n.LocalizeConfig{
			MessageID:    e.BizMessage(),
			TemplateData: e.BizExtra(),
		})
	default:
		errmsg = causeErr.Error()
		errcode = 500
	}

	c.JSON(code, ErrResponse{
		Code: errcode,
		Msg:  errmsg,
	})

}

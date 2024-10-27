package utils

import (
	"backend/library/tools"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	tools.SendErrResponse(ctx, c, code, err)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	tools.SendSuccessResponse(ctx, c, code, data)
}

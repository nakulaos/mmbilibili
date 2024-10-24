package tools

import (
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserID(c *app.RequestContext) int64 {
	ssid, _ := c.Get("uid")
	uid, _ := ssid.(int64)
	return uid
}

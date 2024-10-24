package initializer

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

type TraceInitializer struct {
	EndPoint string
	H        *server.Hertz
}

package initializer

import (
	"backend/app/common/constant"
	"context"
	"github.com/BurntSushi/toml"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	hertzI18n "github.com/hertz-contrib/i18n"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/pprof"
	"golang.org/x/text/language"
)

type MiddlewareInitializer struct {
	h *server.Hertz
}

func NewMiddlewareInitializer(h *server.Hertz) *MiddlewareInitializer {
	return &MiddlewareInitializer{h: h}
}

func (m *MiddlewareInitializer) Init() error {
	m.h.Use(recovery.Recovery())                // recovery
	pprof.Register(m.h)                         // pprof
	m.h.Use(cors.Default())                     // cors
	m.h.Use(gzip.Gzip(gzip.DefaultCompression)) // gzip
	m.h.Use(accesslog.New())                    // access log
	m.h.Use(hertzI18n.Localize(hertzI18n.WithBundle(&hertzI18n.BundleCfg{
		DefaultLanguage:  language.Chinese,
		FormatBundleFile: "toml",
		UnmarshalFunc:    toml.Unmarshal,
		RootPath:         constant.I18nRootDir,
		AcceptLanguage: []language.Tag{
			language.Chinese,
			language.English,
		},
	}),
		hertzI18n.WithGetLangHandle(func(c context.Context, ctx *app.RequestContext, defaultLang string) string {
			lang := ctx.Request.Header.Get("Accept-Language")
			if lang != "" {
				return lang
			}
			lang = ctx.Query("lang")
			if lang == "" {
				return defaultLang
			}
			return lang
		}),
	))

	return nil
}

func MustInitMiddleware(h *server.Hertz) {
	m := NewMiddlewareInitializer(h)
	if err := m.Init(); err != nil {
		panic(err)
	}
}

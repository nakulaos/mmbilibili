package {{.PkgName}}

import (
	"backend/pkg/base"
	"backend/pkg/xvalidator"
	"net/http"


	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

        e := xvalidator.Validate.StructCtx(r.Context(), &req)
        if e != nil {
            base.HttpResult(r, w, nil, e)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			base.HttpResult(r,w,nil,err)
		} else {
			{{if .HasResp}}base.HttpResult(r,w,resp,nil){{else}}base.HttpResult(r,w,nil,nil){{end}}
		}
	}
}

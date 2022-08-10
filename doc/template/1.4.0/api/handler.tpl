package {{.PkgName}}

import (
	"net/http"

    {{if .HasRequest}}"ark-zero-admin/pkg/errorx"{{end}}
	"ark-zero-admin/pkg/response"
	{{.ImportPackages}}
    {{if .HasRequest}}"github.com/go-playground/validator/v10"{{end}}
    "github.com/zeromicro/go-zero/rest/httpx"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
            httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
		    return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.Error(w, err)
			return
		}
        {{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}
	}
}
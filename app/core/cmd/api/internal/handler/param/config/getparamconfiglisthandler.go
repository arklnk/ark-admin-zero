package config

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/param/config"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetParamConfigListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := config.NewGetParamConfigListLogic(r.Context(), svcCtx)
		resp, err := l.GetParamConfigList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

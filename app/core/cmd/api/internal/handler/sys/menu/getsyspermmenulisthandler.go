package menu

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/menu"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysPermMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetSysPermMenuListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysPermMenuList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

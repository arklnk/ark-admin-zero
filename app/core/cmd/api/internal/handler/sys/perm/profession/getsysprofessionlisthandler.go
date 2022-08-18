package profession

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/perm/profession"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysProfessionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := profession.NewGetSysProfessionListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysProfessionList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

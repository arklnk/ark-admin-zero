package dept

import (
	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/dept"
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysDeptListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewGetSysDeptListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysDeptList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

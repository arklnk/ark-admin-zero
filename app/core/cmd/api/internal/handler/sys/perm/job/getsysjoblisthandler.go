package job

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/perm/job"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysJobListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := job.NewGetSysJobListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysJobList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

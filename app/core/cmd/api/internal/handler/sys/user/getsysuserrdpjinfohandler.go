package user

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/user"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysUserRdpjInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetSysUserRdpjInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSysUserRdpjInfo()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

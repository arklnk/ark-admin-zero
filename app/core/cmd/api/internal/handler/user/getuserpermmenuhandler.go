package user

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/user"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserPermMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserPermMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPermMenu()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

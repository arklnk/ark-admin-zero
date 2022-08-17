package user

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/sys/perm/user"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetSysUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysUserList()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

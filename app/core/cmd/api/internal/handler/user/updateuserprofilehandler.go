package user

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/user"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		err := l.UpdateUserProfile()
		if err != nil {
			httpx.Error(w, err)
			return
		}
		response.Response(w, nil, err)
	}
}

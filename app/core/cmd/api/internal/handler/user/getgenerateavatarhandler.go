package user

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/user"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGenerateAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetGenerateAvatarLogic(r.Context(), svcCtx)
		resp, err := l.GetGenerateAvatar()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

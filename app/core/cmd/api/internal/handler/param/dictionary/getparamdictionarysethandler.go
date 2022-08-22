package dictionary

import (
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/logic/param/dictionary"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetParamDictionarySetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dictionary.NewGetParamDictionarySetLogic(r.Context(), svcCtx)
		resp, err := l.GetParamDictionarySet()
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, resp, err)
	}
}

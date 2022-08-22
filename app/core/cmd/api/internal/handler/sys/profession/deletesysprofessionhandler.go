package profession

import (
	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/profession"
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/response"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteSysProfessionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSysProfessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		l := profession.NewDeleteSysProfessionLogic(r.Context(), svcCtx)
		err := l.DeleteSysProfession(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, nil, err)
	}
}

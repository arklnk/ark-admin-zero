package dept

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/sys/perm/dept"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/response"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TransferSysDeptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TransferSysDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		l := dept.NewTransferSysDeptLogic(r.Context(), svcCtx)
		err := l.TransferSysDept(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, nil, err)
	}
}

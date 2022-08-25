package dept

import (
	"errors"
	"net/http"
	"reflect"

	"ark-admin-zero/app/core/cmd/api/internal/logic/sys/dept"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/response"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSysDeptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSysDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		validate := validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		validateErr := translations.RegisterDefaultTranslations(validate, trans)
		if validateErr = validate.StructCtx(r.Context(), req); validateErr != nil {
			for _, err := range validateErr.(validator.ValidationErrors) {
				httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, errors.New(err.Translate(trans)).Error()))
				return
			}
		}

		l := dept.NewUpdateSysDeptLogic(r.Context(), svcCtx)
		err := l.UpdateSysDept(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		response.Response(w, nil, err)
	}
}

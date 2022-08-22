package dictionary

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddParamDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddParamDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddParamDictionaryLogic {
	return &AddParamDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddParamDictionaryLogic) AddParamDictionary(req *types.AddParamDictionaryReq) error {
	_, err := l.svcCtx.SysDictionaryModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err == model.ErrNotFound {
		var dictionary = new(model.SysDictionary)
		err = copier.Copy(dictionary, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		_, err = l.svcCtx.SysDictionaryModel.Insert(l.ctx, dictionary)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddDictionaryErrorCode)
	}
}

package dict

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigDictLogic {
	return &AddConfigDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigDictLogic) AddConfigDict(req *types.AddConfigDictReq) error {
	if req.ParentId != config.SysTopParentId {
		_, err := l.svcCtx.SysDictionaryModel.FindOne(l.ctx, req.ParentId)
		if err != nil {
			return errorx.NewDefaultError(errorx.ParentDictionaryIdErrorCode)
		}
	}
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

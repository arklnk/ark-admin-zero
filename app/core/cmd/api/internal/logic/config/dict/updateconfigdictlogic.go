package dict

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigDictLogic {
	return &UpdateConfigDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigDictLogic) UpdateConfigDict(req *types.UpdateConfigDictReq) error {
	if req.ParentId != config.SysTopParentId {
		_, err := l.svcCtx.SysDictionaryModel.FindOne(l.ctx, req.ParentId)
		if err != nil {
			return errorx.NewDefaultError(errorx.ParentDictionaryIdErrorCode)
		}
	}

	configDictionary, err := l.svcCtx.SysDictionaryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.DictionaryIdErrorCode)
	}

	err = copier.Copy(configDictionary, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	err = l.svcCtx.SysDictionaryModel.Update(l.ctx, configDictionary)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

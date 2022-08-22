package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddParamConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddParamConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddParamConfigLogic {
	return &AddParamConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddParamConfigLogic) AddParamConfig(req *types.AddParamConfigReq) error {
	_, err := l.svcCtx.SysConfigModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err == model.ErrNotFound {
		var config = new(model.SysConfig)
		err = copier.Copy(config, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		_, err = l.svcCtx.SysConfigModel.Insert(l.ctx, config)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddConfigErrorCode)
	}
}

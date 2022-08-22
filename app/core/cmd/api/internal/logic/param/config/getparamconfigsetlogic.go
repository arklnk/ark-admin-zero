package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetParamConfigSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParamConfigSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParamConfigSetLogic {
	return &GetParamConfigSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetParamConfigSetLogic) GetParamConfigSet() (resp *types.ParamConfigSetResp, err error) {
	paramConfigList, err := l.svcCtx.SysConfigModel.FindList(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var config types.ParamConfig
	configList := make([]types.ParamConfig, 0)
	for _, v := range paramConfigList {
		err := copier.Copy(&config, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		configList = append(configList, config)
	}

	return &types.ParamConfigSetResp{
		ConfigList: configList,
	}, nil
}

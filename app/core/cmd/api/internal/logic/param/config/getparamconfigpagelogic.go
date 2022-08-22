package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetParamConfigPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParamConfigPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParamConfigPageLogic {
	return &GetParamConfigPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetParamConfigPageLogic) GetParamConfigPage(req *types.ParamConfigPageReq) (resp *types.ParamConfigPageResp, err error) {
	paramConfigList, err := l.svcCtx.SysConfigModel.FindPageByParentId(l.ctx, req.ParentId, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var config types.ParamConfig
	configList := make([]types.ParamConfig, 0)
	for _, sysJob := range paramConfigList {
		err := copier.Copy(&config, &sysJob)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		configList = append(configList, config)
	}

	total, err := l.svcCtx.SysConfigModel.FindCountByParentId(l.ctx, req.ParentId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.ParamConfigPagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.ParamConfigPageResp{
		ParamConfigList: configList,
		Pagination:      pagination,
	}, nil
}

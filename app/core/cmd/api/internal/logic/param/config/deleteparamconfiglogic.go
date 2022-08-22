package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteParamConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteParamConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteParamConfigLogic {
	return &DeleteParamConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteParamConfigLogic) DeleteParamConfig(req *types.DeleteParamConfigReq) error {
	if req.Id < globalkey.SysMaxConfigId {
		return errorx.NewDefaultError(errorx.NotPermMenuErrorCode)
	}

	err := l.svcCtx.SysConfigModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return nil
}

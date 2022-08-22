package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return nil
}

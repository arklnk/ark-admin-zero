package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateParamConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateParamConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateParamConfigLogic {
	return &UpdateParamConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateParamConfigLogic) UpdateParamConfig(req *types.UpdateParamConfigReq) error {
	// todo: add your logic here and delete this line

	return nil
}

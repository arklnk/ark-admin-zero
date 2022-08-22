package config

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}

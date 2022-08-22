package dictionary

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteParamDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteParamDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteParamDictionaryLogic {
	return &DeleteParamDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteParamDictionaryLogic) DeleteParamDictionary(req *types.DeleteParamDictionaryReq) error {
	// todo: add your logic here and delete this line

	return nil
}

package dictionary

import (
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
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
	if req.Id < globalkey.SysMaxDictionaryId {
		return errorx.NewDefaultError(errorx.NotPermMenuErrorCode)
	}

	err := l.svcCtx.SysDictionaryModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

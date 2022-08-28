package dict

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/config"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigDictLogic {
	return &DeleteConfigDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigDictLogic) DeleteConfigDict(req *types.DeleteConfigDictReq) error {
	if req.Id <= config.SysProtectDictionaryMaxId {
		return errorx.NewDefaultError(errorx.ForbiddenErrorCode)
	}

	total, err := l.svcCtx.SysDictionaryModel.FindCountByParentId(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	if total > 0 {
		return errorx.NewDefaultError(errorx.DeleteDictionaryErrorCode)
	}

	err = l.svcCtx.SysDictionaryModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

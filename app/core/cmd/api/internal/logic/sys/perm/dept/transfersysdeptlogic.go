package dept

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type TransferSysDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferSysDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferSysDeptLogic {
	return &TransferSysDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferSysDeptLogic) TransferSysDept(req *types.TransferSysDeptReq) error {
	sysDept, err := l.svcCtx.SysDeptModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(sysDept, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = l.svcCtx.SysDeptModel.Update(l.ctx, sysDept)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

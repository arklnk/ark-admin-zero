package dept

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDeptLogic {
	return &UpdateSysDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDeptLogic) UpdateSysDept(req *types.UpdateSysDeptReq) error {
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

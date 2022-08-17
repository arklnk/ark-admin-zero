package dept

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDeptLogic {
	return &DeleteSysDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDeptLogic) DeleteSysDept(req *types.DeleteSysDeptReq) error {
	deptList, _ := l.svcCtx.SysDeptModel.FindSubDept(l.ctx, req.Id)
	if len(deptList) == 0 {
		err := l.svcCtx.SysDeptModel.Delete(l.ctx, req.Id)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.DeleteDeptErrorCode)
	}
}

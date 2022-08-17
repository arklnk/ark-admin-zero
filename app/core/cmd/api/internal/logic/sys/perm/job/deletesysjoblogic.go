package job

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysJobLogic {
	return &DeleteSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysJobLogic) DeleteSysJob(req *types.DeleteSysJobReq) error {
	userList, _ := l.svcCtx.SysUserModel.FindByJobId(l.ctx, req.Id)
	if len(userList) == 0 {
		err := l.svcCtx.SysJobModel.Delete(l.ctx, req.Id)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.DeleteJobErrorCode)
	}
}

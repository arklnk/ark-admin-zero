package job

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysJobLogic {
	return &UpdateSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysJobLogic) UpdateSysJob(req *types.UpdateSysJobReq) error {
	sysJob, err := l.svcCtx.SysJobModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.JobIdErrorCode)
	}

	if req.Status == globalkey.SysDisable {
		count, _ := l.svcCtx.SysUserModel.FindCountByJobId(l.ctx, req.Id)
		if count > 0 {
			return errorx.NewDefaultError(errorx.JobIsUsingErrorCode)
		}
	}

	err = copier.Copy(sysJob, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	err = l.svcCtx.SysJobModel.Update(l.ctx, sysJob)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

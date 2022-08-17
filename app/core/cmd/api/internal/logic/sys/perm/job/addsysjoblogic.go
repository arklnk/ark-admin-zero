package job

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysJobLogic {
	return &AddSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysJobLogic) AddSysJob(req *types.AddSysJobReq) error {
	_, err := l.svcCtx.SysJobModel.FindOneByName(l.ctx, req.Name)
	if err == model.ErrNotFound {
		var sysJob = new(model.SysJob)
		err = copier.Copy(sysJob, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		_, err = l.svcCtx.SysJobModel.Insert(l.ctx, sysJob)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddJobErrorCode)
	}
}

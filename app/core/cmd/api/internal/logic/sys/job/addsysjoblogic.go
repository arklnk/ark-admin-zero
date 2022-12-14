package job

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"

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
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		_, err = l.svcCtx.SysJobModel.Insert(l.ctx, sysJob)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddJobErrorCode)
	}
}

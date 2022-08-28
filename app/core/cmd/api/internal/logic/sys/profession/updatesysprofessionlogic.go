package profession

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysProfessionLogic {
	return &UpdateSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysProfessionLogic) UpdateSysProfession(req *types.UpdateSysProfessionReq) error {
	sysProfession, err := l.svcCtx.SysProfessionModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ProfessionIdErrorCode)
	}

	if req.Status == config.SysDisable {
		count, _ := l.svcCtx.SysUserModel.FindCountByProfessionId(l.ctx, req.Id)
		if count > 0 {
			return errorx.NewDefaultError(errorx.JobIsUsingErrorCode)
		}
	}

	err = copier.Copy(sysProfession, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = l.svcCtx.SysProfessionModel.Update(l.ctx, sysProfession)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

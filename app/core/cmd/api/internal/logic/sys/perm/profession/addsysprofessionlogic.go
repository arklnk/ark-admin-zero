package profession

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysProfessionLogic {
	return &AddSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysProfessionLogic) AddSysProfession(req *types.AddSysProfessionReq) error {
	_, err := l.svcCtx.SysProfessionModel.FindOneByName(l.ctx, req.Name)
	if err == model.ErrNotFound {
		var sysProfession = new(model.SysProfession)
		err = copier.Copy(sysProfession, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		_, err = l.svcCtx.SysProfessionModel.Insert(l.ctx, sysProfession)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddProfessionErrorCode)
	}
}

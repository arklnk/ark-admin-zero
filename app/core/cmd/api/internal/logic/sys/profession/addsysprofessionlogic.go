package profession

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"

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
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		_, err = l.svcCtx.SysProfessionModel.Insert(l.ctx, sysProfession)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddProfessionErrorCode)
	}
}

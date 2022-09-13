package dept

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDeptLogic {
	return &AddSysDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysDeptLogic) AddSysDept(req *types.AddSysDeptReq) error {
	_, err := l.svcCtx.SysDeptModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err == model.ErrNotFound {
		if req.ParentId != config.SysTopParentId {
			_, err := l.svcCtx.SysDeptModel.FindOne(l.ctx,req.ParentId)
			if err != nil {
				return errorx.NewDefaultError(errorx.ParentDeptIdErrorCode)
			}
		}

		var sysDept = new(model.SysDept)
		err = copier.Copy(sysDept, req)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		_, err = l.svcCtx.SysDeptModel.Insert(l.ctx, sysDept)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		return nil
	} else {
		return errorx.NewDefaultError(errorx.AddDeptErrorCode)
	}
}

package dept

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

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
	if req.ParentId != config.SysTopParentId {
		_, err := l.svcCtx.SysDeptModel.FindOne(l.ctx,req.ParentId)
		if err != nil {
			return errorx.NewDefaultError(errorx.ParentDeptIdErrorCode)
		}
	}

	if req.Id == req.ParentId {
		return errorx.NewDefaultError(errorx.ParentDeptErrorCode)
	}

	dept, err := l.svcCtx.SysDeptModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err != model.ErrNotFound && dept.Id != req.Id {
		return errorx.NewDefaultError(errorx.UpdateDeptUniqueKeyErrorCode)
	}

	deptIds := make([]uint64, 0)
	deptIds = l.getSubDept(deptIds, req.Id)
	if utils.ArrayContainValue(deptIds, req.ParentId) {
		return errorx.NewDefaultError(errorx.SetParentIdErrorCode)
	}

	sysDept, err := l.svcCtx.SysDeptModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.DeptIdErrorCode)
	}

	err = copier.Copy(sysDept, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	err = l.svcCtx.SysDeptModel.Update(l.ctx, sysDept)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

func (l *UpdateSysDeptLogic) getSubDept(deptIds []uint64, id uint64) []uint64 {
	deptList, err := l.svcCtx.SysDeptModel.FindSubDept(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return deptIds
	}

	for _, v := range deptList {
		deptIds = append(deptIds, v.Id)
		deptIds = l.getSubDept(deptIds, v.Id)
	}

	return deptIds
}

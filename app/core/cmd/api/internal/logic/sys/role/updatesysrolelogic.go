package role

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysRoleLogic {
	return &UpdateSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysRoleLogic) UpdateSysRole(req *types.UpdateSysRoleReq) error {
	if req.ParentId != config.SysTopParentId {
		_, err := l.svcCtx.SysRoleModel.FindOne(l.ctx,req.ParentId)
		if err != nil {
			return errorx.NewDefaultError(errorx.ParentRoleIdErrorCode)
		}
	}

	if req.Id == config.SysSuperRoleId {
		return errorx.NewDefaultError(errorx.NotPermMenuErrorCode)
	}

	if req.Id == req.ParentId {
		return errorx.NewDefaultError(errorx.ParentRoleErrorCode)
	}

	role, err := l.svcCtx.SysRoleModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err != model.ErrNotFound && role.Id != req.Id {
		return errorx.NewDefaultError(errorx.UpdateRoleUniqueKeyErrorCode)
	}

	roleIds := make([]uint64, 0)
	roleIds = l.getSubRole(roleIds, req.Id)
	if utils.ArrayContainValue(roleIds, req.ParentId) {
		return errorx.NewDefaultError(errorx.SetParentIdErrorCode)
	}

	sysRole, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.RoleIdErrorCode)
	}

	err = copier.Copy(sysRole, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	bytes, err := json.Marshal(req.PermMenuIds)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	sysRole.PermMenuIds = string(bytes)
	err = l.svcCtx.SysRoleModel.Update(l.ctx, sysRole)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

func (l *UpdateSysRoleLogic) getSubRole(roleIds []uint64, id uint64) []uint64 {
	roleList, err := l.svcCtx.SysRoleModel.FindSubRole(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return roleIds
	}

	for _, v := range roleList {
		roleIds = append(roleIds, v.Id)
		roleIds = l.getSubRole(roleIds, v.Id)
	}

	return roleIds
}

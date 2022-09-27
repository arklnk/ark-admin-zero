package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserLogic {
	return &UpdateSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserLogic) UpdateSysUser(req *types.UpdateSysUserReq) error {
	currentUserId := utils.GetUserId(l.ctx)
	var currentUserRoleIds []int64
	var roleIds []int64
	if currentUserId == config.SysSuperUserId {
		sysRoleList, _ := l.svcCtx.SysRoleModel.FindAll(l.ctx)
		for _, role := range sysRoleList {
			currentUserRoleIds=append(currentUserRoleIds,role.Id)
			roleIds=append(roleIds,role.Id)
		}
		
	}else {
		currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
		err := json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRoleIds)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		
		roleIds = append(roleIds, currentUserRoleIds...)
	}

	editUser, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	var editUserRoleIds []int64
	err = json.Unmarshal([]byte(editUser.RoleIds), &editUserRoleIds)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	roleIds = append(roleIds, editUserRoleIds...)
	
	for _, id := range req.RoleIds {
		if !utils.ArrayContainValue(roleIds, id) {
			return errorx.NewDefaultError(errorx.AssigningRolesErrorCode)
		}
	}

	for _, id := range utils.Difference(editUserRoleIds, currentUserRoleIds) {
		if !utils.ArrayContainValue(req.RoleIds, id) {
			return errorx.NewDefaultError(errorx.AssigningRolesErrorCode)
		}
	}

	_, err = l.svcCtx.SysDeptModel.FindOne(l.ctx, req.DeptId)
	if err != nil {
		return errorx.NewDefaultError(errorx.DeptIdErrorCode)
	}

	_, err = l.svcCtx.SysProfessionModel.FindOne(l.ctx, req.ProfessionId)
	if err != nil {
		return errorx.NewDefaultError(errorx.ProfessionIdErrorCode)
	}

	_, err = l.svcCtx.SysJobModel.FindOne(l.ctx, req.JobId)
	if err != nil {
		return errorx.NewDefaultError(errorx.JobIdErrorCode)
	}

	err = copier.Copy(editUser, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	bytes, err := json.Marshal(req.RoleIds)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.Redis.Del(config.SysPermMenuCachePrefix + strconv.FormatInt(editUser.Id, 10))
	_, err = l.svcCtx.Redis.Del(config.SysOnlineUserCachePrefix + strconv.FormatInt(editUser.Id, 10))
	editUser.RoleIds = string(bytes)
	err = l.svcCtx.SysUserModel.Update(l.ctx, editUser)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

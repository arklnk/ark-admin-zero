package menu

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysPermMenuLogic {
	return &DeleteSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysPermMenuLogic) DeleteSysPermMenu(req *types.DeleteSysPermMenuReq) error {
	currentUserId := utils.GetUserId(l.ctx)
	if currentUserId != config.SysSuperUserId {
		var currentUserPermMenuIds []int64
		currentUserPermMenuIds = l.getCurrentUserPermMenuIds(currentUserId)
		if !utils.ArrayContainValue(currentUserPermMenuIds, req.Id) {
			return errorx.NewDefaultError(errorx.NotPermMenuErrorCode)
		}
	}

	if req.Id <= config.SysProtectPermMenuMaxId {
		return errorx.NewDefaultError(errorx.ForbiddenErrorCode)
	}

	count, _ := l.svcCtx.SysPermMenuModel.FindCountByParentId(l.ctx, req.Id)
	if count != 0 {
		return errorx.NewDefaultError(errorx.DeletePermMenuErrorCode)
	}

	err := l.svcCtx.SysPermMenuModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

func (l *DeleteSysPermMenuLogic) getCurrentUserPermMenuIds(currentUserId int64) (ids []int64) {
	var currentPermMenuIds []int64
	if currentUserId != config.SysSuperUserId {
		var currentUserRoleIds []int64
		var roleIds []int64
		currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
		_ = json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRoleIds)
		roleIds = append(roleIds, currentUserRoleIds...)
		var ids string
		for i, v := range roleIds {
			if i == 0 {
				ids = strconv.FormatInt(v, 10)
			}
			ids = ids + "," + strconv.FormatInt(v, 10)
		}

		sysRoles, _ := l.svcCtx.SysRoleModel.FindByIds(l.ctx, ids)
		var rolePermMenus []int64
		for _, v := range sysRoles {
			err := json.Unmarshal([]byte(v.PermMenuIds), &rolePermMenus)
			if err != nil {
				return nil
			}
			currentPermMenuIds = append(currentPermMenuIds, rolePermMenus...)
		}
	}

	return currentPermMenuIds
}

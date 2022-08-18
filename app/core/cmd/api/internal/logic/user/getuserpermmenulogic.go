package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
	"ark-admin-zero/common/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermMenuLogic {
	return &GetUserPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermMenuLogic) GetUserPermMenu() (resp *types.UserPermMenuResp, err error) {
	userId := utils.GetUserId(l.ctx)

	// 查询用户信息
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var roles []int64
	// 用户所属角色
	err = json.Unmarshal([]byte(user.RoleIds), &roles)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var permMenu []int64
	var userPermMenu []*model.SysPermMenu
	userPermMenu, permMenu, err = l.countUserPermMenu(roles, permMenu)
	var menus []types.Menu
	var perms []string
	if err != nil {
		return &types.UserPermMenuResp{Menus: menus, Perms: perms}, nil
	}

	for _, v := range userPermMenu {
		var menu types.Menu
		err := copier.Copy(menu, v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		menus = append(menus, menu)
		var permArr []string
		err = json.Unmarshal([]byte(v.Perms), &permArr)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		for _, p := range permArr {
			p = globalkey.PermMenuPrefix + p
			_, err := l.svcCtx.Redis.Sadd(globalkey.CachePermMenuKey+strconv.FormatInt(userId, 10), p)
			if err != nil {
				return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
			}
			perms = append(perms, p)
		}

	}

	return &types.UserPermMenuResp{Menus: menus, Perms: utils.ArrayUniqueValue[string](perms)}, nil
}

func (l *GetUserPermMenuLogic) countUserPermMenu(roles []int64, permMenu []int64) ([]*model.SysPermMenu, []int64, error) {
	if utils.ArrayContainValue(roles, globalkey.SuperAdminRole) {
		sysPermMenus, err := l.svcCtx.SysPermMenuModel.FindAll(l.ctx)
		if err != nil {
			return nil, permMenu, err
		}

		return sysPermMenus, permMenu, nil
	} else {
		for _, roleId := range roles {
			// 查询角色信息
			role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, roleId)
			if err != nil && err != model.ErrNotFound {
				return nil, permMenu, errorx.NewDefaultError(errorx.ServerErrorCode)
			}

			var perms []int64
			// 角色所拥有的权限id
			err = json.Unmarshal([]byte(role.PermMenuIds), &perms)
			if err != nil {
				return nil, permMenu, errorx.NewDefaultError(errorx.ServerErrorCode)
			}

			// 汇总用户所属角色权限id
			permMenu = append(permMenu, perms...)
			permMenu = l.getRolePermMenu(permMenu, roleId)
		}

		// 过滤重复的权限id
		permMenu = utils.ArrayUniqueValue[int64](permMenu)
		var roleIds string
		for i, id := range permMenu {
			if i == 0 {
				roleIds = strconv.FormatInt(id, 10)
				continue
			}
			roleIds = roleIds + "," + strconv.FormatInt(id, 10)
		}

		// 根据权限id获取具体权限
		sysPermMenus, err := l.svcCtx.SysPermMenuModel.FindByIds(l.ctx, roleIds)
		if err != nil {
			return nil, permMenu, err
		}

		return sysPermMenus, permMenu, nil
	}
}

func (l *GetUserPermMenuLogic) getRolePermMenu(perms []int64, roleId int64) []int64 {
	roles, err := l.svcCtx.SysRoleModel.FindSubRole(l.ctx, roleId)
	if err != nil && err != model.ErrNotFound {
		return perms
	}

	for _, role := range roles {
		var subPerms []int64
		err = json.Unmarshal([]byte(role.PermMenuIds), &subPerms)
		perms = append(perms, subPerms...)
		perms = l.getRolePermMenu(perms, role.Id)
	}

	return perms
}

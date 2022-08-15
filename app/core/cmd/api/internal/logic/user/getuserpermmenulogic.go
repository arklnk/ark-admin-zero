package user

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/sysconstant"
	"ark-zero-admin/common/utils"

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

func (l *GetUserPermMenuLogic) GetUserPermMenu() (resp *types.PermMenuResp, err error) {
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
	for _, roleId := range roles {

		// 查询角色信息
		role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, roleId)
		if err != nil && err != model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		var perms []int64

		// 角色所拥有的权限id
		err = json.Unmarshal([]byte(role.PermMenuIds), &perms)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		// 汇总用户所属角色权限id
		permMenu = append(permMenu, perms...)
		permMenu = l.getSubRolePermMenu(permMenu, roleId)
	}

	// 过滤重复的权限id
	permMenu = utils.ArrayUniqueValue[int64](permMenu)

	roleIds := "0"
	for _, id := range permMenu {
		roleIds = roleIds + "," + strconv.FormatInt(id, 10)
	}

	// 根据权限id获取具体权限
	userPermMenu, err := l.svcCtx.SysPermMenuModel.FindUserPermMenu(l.ctx, roleIds)
	var menus []types.Menu
	var perms []string
	if err != nil {
		return &types.PermMenuResp{Menus: menus, Perms: perms}, nil
	}
	for _, perm := range userPermMenu {
		var menu types.Menu
		err := copier.Copy(&menu, perm)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		menus = append(menus, menu)
		arr := strings.Split(perm.Perms, ";")
		for _, s := range arr {
			_, err := l.svcCtx.Redis.Sadd(sysconstant.CachePermMenuKey+strconv.FormatInt(userId, 10), s)
			if err != nil {
				return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
			}
			perms = append(perms, s)
		}
	}

	return &types.PermMenuResp{Menus: menus, Perms: utils.ArrayUniqueValue[string](perms)}, nil
}

func (l *GetUserPermMenuLogic) getSubRolePermMenu(perms []int64, roleId int64) []int64 {
	roles, err := l.svcCtx.SysRoleModel.FindSubRole(l.ctx, roleId)
	if err != nil && err != model.ErrNotFound {
		return perms
	}
	for _, role := range roles {
		var subPerms []int64
		err = json.Unmarshal([]byte(role.PermMenuIds), &subPerms)
		perms = append(perms, subPerms...)
		perms = l.getSubRolePermMenu(perms, role.Id)
	}
	return perms
}

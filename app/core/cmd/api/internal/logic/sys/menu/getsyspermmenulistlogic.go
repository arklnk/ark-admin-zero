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

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysPermMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermMenuListLogic {
	return &GetSysPermMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysPermMenuListLogic) GetSysPermMenuList() (resp *types.SysPermMenuListResp, err error) {
	permMenus, err := l.svcCtx.SysPermMenuModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	currentUserId := utils.GetUserId(l.ctx)
	var currentUserPermMenuIds []uint64
	if currentUserId != config.SysProtectUserId {
		currentUserPermMenuIds = l.getCurrentUserPermMenuIds(currentUserId)
	}

	var menu types.PermMenu
	PermMenuList := make([]types.PermMenu, 0)
	for _, v := range permMenus {
		err := copier.Copy(&menu, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		var perms []string
		err = json.Unmarshal([]byte(v.Perms), &perms)
		menu.Perms = perms
		if currentUserId == config.SysProtectUserId {
			menu.Has = 1
		} else {
			if utils.ArrayContainValue(currentUserPermMenuIds, v.Id) {
				menu.Has = 1
			} else {
				menu.Has = 0
			}
		}
		PermMenuList = append(PermMenuList, menu)
	}

	return &types.SysPermMenuListResp{PermMenuList: PermMenuList}, nil
}

func (l *GetSysPermMenuListLogic) getCurrentUserPermMenuIds(currentUserId uint64) (ids []uint64) {
	var currentPermMenuIds []uint64
	if currentUserId != config.SysProtectUserId {
		var currentUserRoleIds []uint64
		var roleIds []uint64
		currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
		_ = json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRoleIds)
		roleIds = append(roleIds, currentUserRoleIds...)
		var ids string
		for i, v := range roleIds {
			if i == 0 {
				ids = strconv.FormatUint(v, 10)
			}
			ids = ids + "," + strconv.FormatUint(v, 10)
		}

		sysRoles, _ := l.svcCtx.SysRoleModel.FindByIds(l.ctx, ids)
		var rolePermMenus []uint64
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

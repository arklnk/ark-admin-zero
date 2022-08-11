package user

import (
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/pkg/utils"
	"context"
	"encoding/json"
	"fmt"

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
	userId := utils.UserId(l.ctx)
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	var roles []int64
	err = json.Unmarshal([]byte(user.RoleIds), &roles)
	if err != nil {
		return nil, err
	}
	var permMenu []int64
	for _, roleId := range roles {
		role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, roleId)
		if err != nil && err != model.ErrNotFound {
			return nil, err
		}
		var perms []int64
		err = json.Unmarshal([]byte(role.PermMenuIds), &perms)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%v", role.PermMenuIds)
		permMenu = append(permMenu, perms...)
		//l.SubRoleCallback(roleId)
	}
	fmt.Printf("%v", permMenu)
	return
}

func (l *GetUserPermMenuLogic) SubRoleCallback(roleId int64) {
	roles, err := l.svcCtx.SysRoleModel.FindSubRole(l.ctx, roleId)
	if err != nil && err != model.ErrNotFound {
		return
	}
	for _, role := range roles {
		fmt.Printf("%v", role)
		l.SubRoleCallback(role.Id)
	}
}

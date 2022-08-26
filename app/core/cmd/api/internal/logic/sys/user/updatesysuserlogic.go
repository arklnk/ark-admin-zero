package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"

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
	currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
	var currentUserRole []uint64
	err := json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRole)
	if err != nil {
		return nil
	}

	var roleIds []uint64
	roleIds = append(roleIds, currentUserRole...)

	editUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	var editUserRole []uint64
	err = json.Unmarshal([]byte(editUser.RoleIds), &editUserRole)
	if err != nil {
		return nil
	}
	roleIds = append(roleIds, editUserRole...)

	for _, id := range req.RoleIds {
		if !utils.ArrayContainValue(roleIds, id) {
			return errorx.NewDefaultError(errorx.AssigningRolesErrorCode)
		}
	}

	for _, id := range utils.Difference(editUserRole, currentUserRole) {
		if !utils.ArrayContainValue(req.RoleIds, id) {
			return errorx.NewDefaultError(errorx.AssigningRolesErrorCode)
		}
	}

	err = copier.Copy(editUser, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	bytes, err := json.Marshal(req.RoleIds)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	_, err = l.svcCtx.Redis.Del(config.SysPermMenuCachePrefix + strconv.FormatUint(editUser.Id, 10))
	_, err = l.svcCtx.Redis.Del(config.SysOnlineUserCachePrefix + strconv.FormatUint(editUser.Id, 10))
	editUser.RoleIds = string(bytes)
	err = l.svcCtx.SysUserModel.Update(l.ctx, editUser)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

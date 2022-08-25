package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"

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
	sysUser, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	err = copier.Copy(sysUser, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	bytes, err := json.Marshal(req.RoleIds)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	_, err = l.svcCtx.Redis.Del(config.SysPermMenuCachePrefix + strconv.FormatUint(sysUser.Id, 10))
	_, err = l.svcCtx.Redis.Del(config.SysOnlineUserCachePrefix + strconv.FormatUint(sysUser.Id, 10))
	sysUser.RoleIds = string(bytes)
	err = l.svcCtx.SysUserModel.Update(l.ctx, sysUser)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

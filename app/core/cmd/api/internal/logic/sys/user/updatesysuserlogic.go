package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"

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

	_, err := l.svcCtx.SysDeptModel.FindOne(l.ctx, req.DeptId)
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

	editUser, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	bytes, err := json.Marshal(req.RoleIds)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	editUser.RoleIds = string(bytes)

	err = l.svcCtx.SysUserModel.Update(l.ctx, editUser)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.Redis.Del(globalkey.SysPermMenuCachePrefix + strconv.FormatInt(editUser.Id, 10))
	_, err = l.svcCtx.Redis.Del(globalkey.SysOnlineUserCachePrefix + strconv.FormatInt(editUser.Id, 10))

	return nil
}

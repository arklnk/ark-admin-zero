package user

import (
	"context"
	"encoding/json"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/utils"

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
		return err
	}

	sysUser.RoleIds = string(bytes)
	sysUser.Birthday = utils.StrToTime(req.Birthday)
	err = l.svcCtx.SysUserModel.Update(l.ctx, sysUser)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

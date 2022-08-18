package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserPasswordLogic {
	return &UpdateSysUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserPasswordLogic) UpdateSysUserPassword(req *types.UpdateSysUserPasswordReq) error {
	sysUser, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(sysUser, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	sysUser.Password = utils.MD5(req.Password + l.svcCtx.Config.Salt)
	err = l.svcCtx.SysUserModel.Update(l.ctx, sysUser)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

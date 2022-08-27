package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"

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
		return errorx.NewDefaultError(errorx.UserIdErrorCode)
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

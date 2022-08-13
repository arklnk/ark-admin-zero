package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/pkg/errorx"
	"ark-zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.PasswordReq) error {
	userId := utils.GetUserId(l.ctx)
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	if user.Password != utils.MD5(req.OldPassword+l.svcCtx.Config.Salt) {
		return errorx.NewDefaultError(errorx.PasswordErrorCode)
	}
	user.Password = utils.MD5(req.NewPassword + l.svcCtx.Config.Salt)
	err = l.svcCtx.SysUserModel.Update(l.ctx, user)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	return nil
}

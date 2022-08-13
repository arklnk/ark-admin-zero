package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/pkg/errorx"
	"ark-zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.UpdateProfileResp) error {
	userId := utils.GetUserId(l.ctx)
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	user.Username = req.Username
	user.Nickname = req.Nickname
	user.Gender = req.Gender
	user.Email = req.Email
	user.Mobile = req.Mobile
	user.Remark = req.Remark
	user.Avatar = req.Avatar
	err = l.svcCtx.SysUserModel.Update(l.ctx, user)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

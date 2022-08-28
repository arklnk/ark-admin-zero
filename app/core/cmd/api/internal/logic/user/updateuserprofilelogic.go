package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
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

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.UpdateProfileReq) error {
	dictionary, err := l.svcCtx.SysDictionaryModel.FindOneByUniqueKey(l.ctx, "sys_userinfo")
	if dictionary.Status == config.SysDisable {
		return errorx.NewDefaultError(errorx.ForbiddenErrorCode)
	}

	userId := utils.GetUserId(l.ctx)
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(user,req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = l.svcCtx.SysUserModel.Update(l.ctx, user)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

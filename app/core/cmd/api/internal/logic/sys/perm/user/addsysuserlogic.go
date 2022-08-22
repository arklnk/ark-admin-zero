package user

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
	"ark-admin-zero/common/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysUserLogic {
	return &AddSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysUserLogic) AddSysUser(req *types.AddSysUserReq) error {
	_, err := l.svcCtx.SysUserModel.FindOneByAccount(l.ctx, req.Account)
	if err == model.ErrNotFound {
		var sysUser = new(model.SysUser)
		err = copier.Copy(sysUser, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		bytes, err := json.Marshal(req.RoleIds)
		if err != nil {
			return err
		}

		sysUser.RoleIds = string(bytes)
		sysUser.Password = utils.MD5(globalkey.SysNewUserDefaultPassword + l.svcCtx.Config.Salt)
		_, err = l.svcCtx.SysUserModel.Insert(l.ctx, sysUser)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddUserErrorCode)
	}
}

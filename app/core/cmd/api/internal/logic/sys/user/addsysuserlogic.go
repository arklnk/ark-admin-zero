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

		var sysUser = new(model.SysUser)
		err = copier.Copy(sysUser, req)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		bytes, err := json.Marshal(req.RoleIds)
		sysUser.RoleIds = string(bytes)
		dictionary, err := l.svcCtx.SysDictionaryModel.FindOneByUniqueKey(l.ctx, "sys_pwd")
		var password string
		if dictionary.Status == globalkey.SysEnable {
			password = dictionary.Value
		} else {
			password = globalkey.SysNewUserDefaultPassword
		}

		sysUser.Password = utils.MD5(password + l.svcCtx.Config.Salt)
		sysUser.Avatar = utils.AvatarUrl()
		_, err = l.svcCtx.SysUserModel.Insert(l.ctx, sysUser)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddUserErrorCode)
	}
}

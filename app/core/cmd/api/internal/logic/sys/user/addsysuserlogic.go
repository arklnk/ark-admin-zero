package user

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"
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
		currentUserId := utils.GetUserId(l.ctx)
		currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
		var currentUserRole []uint64
		err := json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRole)
		if err != nil {
			return nil
		}

		var roleIds []uint64
		roleIds = append(roleIds, currentUserRole...)
		for _, id := range req.RoleIds {
			if !utils.ArrayContainValue(roleIds,id) {
				return errorx.NewDefaultError(errorx.AssigningRolesErrorCode)
			}
		}

		var sysUser = new(model.SysUser)
		err = copier.Copy(sysUser, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		bytes, err := json.Marshal(req.RoleIds)
		sysUser.RoleIds = string(bytes)
		dictionary, err := l.svcCtx.SysDictionaryModel.FindOneByUniqueKey(l.ctx, "sys_pwd")
		var password string
		if dictionary.Status == config.SysEnable {
			password = dictionary.Value
		} else {
			password = config.SysNewUserDefaultPassword
		}

		sysUser.Password = utils.MD5(password + l.svcCtx.Config.Salt)
		sysUser.Avatar = utils.AvatarUrl()
		_, err = l.svcCtx.SysUserModel.Insert(l.ctx, sysUser)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddUserErrorCode)
	}
}

package role

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysRoleLogic {
	return &AddSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysRoleLogic) AddSysRole(req *types.AddSysRoleReq) error {
	_, err := l.svcCtx.SysRoleModel.FindOneByUniqueKey(l.ctx, req.UniqueKey)
	if err == model.ErrNotFound {
		if req.ParentId != globalkey.SysTopParentId {
			_, err := l.svcCtx.SysRoleModel.FindOne(l.ctx,req.ParentId)
			if err != nil {
				return errorx.NewDefaultError(errorx.ParentRoleIdErrorCode)
			}
		}

		var sysRole = new(model.SysRole)
		err = copier.Copy(sysRole, req)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		bytes, err := json.Marshal(req.PermMenuIds)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		sysRole.PermMenuIds = string(bytes)
		_, err = l.svcCtx.SysRoleModel.Insert(l.ctx, sysRole)
		if err != nil {
			return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddRoleErrorCode)
	}
}

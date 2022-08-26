package role

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"

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
		if req.ParentId != config.SysTopParentId {
			_, err := l.svcCtx.SysRoleModel.FindOne(l.ctx,req.ParentId)
			if err != nil {
				return errorx.NewDefaultError(errorx.ParentRoleIdErrorCode)
			}
		}

		var sysRole = new(model.SysRole)
		err = copier.Copy(sysRole, req)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		bytes, err := json.Marshal(req.PermMenuIds)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		sysRole.PermMenuIds = string(bytes)
		_, err = l.svcCtx.SysRoleModel.Insert(l.ctx, sysRole)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.AddRoleErrorCode)
	}
}

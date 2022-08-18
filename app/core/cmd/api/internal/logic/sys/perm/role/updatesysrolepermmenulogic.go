package role

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysRolePermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysRolePermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysRolePermMenuLogic {
	return &UpdateSysRolePermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysRolePermMenuLogic) UpdateSysRolePermMenu(req *types.UpdateSysRolePermMenuReq) error {
	sysRole, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	bytes, err := json.Marshal(req.PermMenuIds)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	sysRole.PermMenuIds = string(bytes)
	err = l.svcCtx.SysRoleModel.Update(l.ctx, sysRole)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

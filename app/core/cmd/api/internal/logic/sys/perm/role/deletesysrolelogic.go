package role

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysRoleLogic {
	return &DeleteSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysRoleLogic) DeleteSysRole(req *types.DeleteSysRoleReq) error {
	roleList, _ := l.svcCtx.SysRoleModel.FindSubRole(l.ctx, req.Id)
	if len(roleList) == 0 {
		err := l.svcCtx.SysRoleModel.Delete(l.ctx, req.Id)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.DeleteRoleErrorCode)
	}
}

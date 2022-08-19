package menu

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysPermMenuLogic {
	return &DeleteSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysPermMenuLogic) DeleteSysPermMenu(req *types.DeleteSysPermMenuReq) error {
	subPermMenu, _ := l.svcCtx.SysPermMenuModel.FindSubPermMenuById(l.ctx, req.Id)
	if len(subPermMenu) != 0 {
		return errorx.NewDefaultError(errorx.DeletePermMenuErrorCode)
	}

	err := l.svcCtx.SysPermMenuModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

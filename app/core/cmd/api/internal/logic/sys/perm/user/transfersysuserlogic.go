package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferSysUserLogic {
	return &TransferSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferSysUserLogic) TransferSysUser(req *types.TransferSysUserReq) error {
	for _, id := range req.Ids {
		sysUser, err := l.svcCtx.SysUserModel.FindOne(l.ctx, id)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		sysUser.DeptId = req.DeptId
		err = l.svcCtx.SysUserModel.Update(l.ctx, sysUser)
	}

	return nil
}

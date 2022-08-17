package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return nil
}

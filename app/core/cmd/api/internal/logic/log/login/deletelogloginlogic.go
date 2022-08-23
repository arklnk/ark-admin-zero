package login

import (
	"ark-admin-zero/common/errorx"
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogLoginLogic {
	return &DeleteLogLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogLoginLogic) DeleteLogLogin(req *types.DeleteLogLoginReq) error {
	err := l.svcCtx.SysLogModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

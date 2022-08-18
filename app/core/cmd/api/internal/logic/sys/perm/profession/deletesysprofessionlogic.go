package profession

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysProfessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysProfessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysProfessionLogic {
	return &DeleteSysProfessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysProfessionLogic) DeleteSysProfession(req *types.DeleteSysProfessionReq) error {
	userList, _ := l.svcCtx.SysUserModel.FindByCondition(l.ctx, "profession_id", req.Id)
	if len(userList) == 0 {
		err := l.svcCtx.SysProfessionModel.Delete(l.ctx, req.Id)
		if err != nil {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		return nil
	} else {

		return errorx.NewDefaultError(errorx.DeleteProfessionErrorCode)
	}
}

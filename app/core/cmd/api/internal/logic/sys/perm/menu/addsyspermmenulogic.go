package menu

import (
	"context"
	"encoding/json"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPermMenuLogic {
	return &AddSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysPermMenuLogic) AddSysPermMenu(req *types.AddSysPermMenuReq) error {
	var permMenu = new(model.SysPermMenu)
	err := copier.Copy(permMenu, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	bytes, err := json.Marshal(req.Perms)
	if err != nil {
		return err
	}
	permMenu.Perms = string(bytes)
	_, err = l.svcCtx.SysPermMenuModel.Insert(l.ctx, permMenu)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	return nil
}

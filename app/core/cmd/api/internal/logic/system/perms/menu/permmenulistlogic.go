package menu

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermMenuListLogic {
	return &PermMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermMenuListLogic) PermMenuList() (resp *types.PermMenuListResp, err error) {
	permMenus, err := l.svcCtx.SysPermMenuModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	var menu types.PermMenu
	var PermMenuList []types.PermMenu
	for _, v := range permMenus {
		menu.Id = v.Id
		menu.ParentId = v.ParentId
		menu.Name = v.Name
		menu.Router = v.Router
		menu.Perms = v.Perms
		menu.Type = v.Type
		menu.Icon = v.Icon
		menu.OrderNum = v.OrderNum
		menu.ViewPath = v.ViewPath
		menu.IsShow = v.IsShow
		menu.ActiveRouter = v.ActiveRouter

		PermMenuList = append(PermMenuList, menu)
	}
	return &types.PermMenuListResp{PermMenuList: PermMenuList}, nil
}

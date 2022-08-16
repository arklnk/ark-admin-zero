package menu

import (
	"context"
	"encoding/json"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysPermMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermMenuListLogic {
	return &GetSysPermMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysPermMenuListLogic) GetSysPermMenuList() (resp *types.SysPermMenuListResp, err error) {
	permMenus, err := l.svcCtx.SysPermMenuModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var menu types.PermMenu
	var PermMenuList []types.PermMenu
	var perms []string
	for _, v := range permMenus {
		err := copier.Copy(&menu, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.PasswordErrorCode)
		}
		err = json.Unmarshal([]byte(v.Perms), &perms)
		menu.Perms = perms
		PermMenuList = append(PermMenuList, menu)
	}

	return &types.SysPermMenuListResp{PermMenuList: PermMenuList}, nil
}

package menu

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
	"ark-admin-zero/common/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPermMenuLogic {
	return &UpdateSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysPermMenuLogic) UpdateSysPermMenu(req *types.UpdateSysPermMenuReq) error {
	if req.Id == req.ParentId {
		return errorx.NewDefaultError(errorx.ParentPermMenuErrorCode)
	}

	if req.ParentId != globalkey.SysTopMenuId {
		parentPermMenu, _ := l.svcCtx.SysPermMenuModel.FindOne(l.ctx, req.ParentId)
		if parentPermMenu.Type == 2 {
			return errorx.NewDefaultError(errorx.SetParentTypeErrorCode)
		}
	}

	permMenuIds := make([]int64, 0)
	permMenuIds = l.getSubPermMenu(permMenuIds, req.Id)
	if utils.ArrayContainValue(permMenuIds, req.ParentId) {
		return errorx.NewDefaultError(errorx.SetParentIdErrorCode)
	}

	permMenu, err := l.svcCtx.SysPermMenuModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(permMenu, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	bytes, err := json.Marshal(req.Perms)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	permMenu.Perms = string(bytes)
	err = l.svcCtx.SysPermMenuModel.Update(l.ctx, permMenu)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}

func (l *UpdateSysPermMenuLogic) getSubPermMenu(permMenuIds []int64, id int64) []int64 {
	permMenuList, err := l.svcCtx.SysPermMenuModel.FindSubPermMenu(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return permMenuIds
	}

	for _, v := range permMenuList {
		permMenuIds = append(permMenuIds, v.Id)
		permMenuIds = l.getSubPermMenu(permMenuIds, v.Id)
	}

	return permMenuIds
}

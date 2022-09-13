package menu

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"
	"ark-admin-zero/config"

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
	userId := utils.GetUserId(l.ctx)
	if userId != config.SysSuperUserId {
		for _, v := range req.Perms {
			is, err := l.svcCtx.Redis.Sismember(config.SysPermMenuCachePrefix+strconv.FormatUint(userId, 10), config.SysPermMenuPrefix+v)
			if err != nil || is != true {
				return errorx.NewDefaultError(errorx.NotPermMenuErrorCode)
			}
		}
	}

	if req.ParentId != config.SysTopParentId {
		parentPermMenu, err := l.svcCtx.SysPermMenuModel.FindOne(l.ctx, req.ParentId)
		if err != nil {
			return errorx.NewDefaultError(errorx.ParentPermMenuIdErrorCode)
		}

		if parentPermMenu.Type == 2 {
			return errorx.NewDefaultError(errorx.SetParentTypeErrorCode)
		}
	}

	var permMenu = new(model.SysPermMenu)
	err := copier.Copy(permMenu, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	bytes, err := json.Marshal(req.Perms)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	permMenu.Perms = string(bytes)
	_, err = l.svcCtx.SysPermMenuModel.Insert(l.ctx, permMenu)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}

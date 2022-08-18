package role

import (
	"context"
	"encoding/json"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysRoleListLogic {
	return &GetSysRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysRoleListLogic) GetSysRoleList() (resp *types.SysRoleListResp, err error) {
	sysRoleList, err := l.svcCtx.SysRoleModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var role types.Role
	var roleList []types.Role
	for _, v := range sysRoleList {
		err := copier.Copy(&role, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		var permMenuIds []int64
		err = json.Unmarshal([]byte(v.PermMenuIds), &permMenuIds)
		role.PermMenuIds = permMenuIds
		roleList = append(roleList, role)
	}

	return &types.SysRoleListResp{
		RoleList: roleList,
	}, nil
}

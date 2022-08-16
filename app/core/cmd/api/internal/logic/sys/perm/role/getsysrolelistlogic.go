package role

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

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
	sysRoles, err := l.svcCtx.SysRoleModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var role types.Role
	var roleList []types.Role
	for _, sysRole := range sysRoles {
		err := copier.Copy(&role, &sysRole)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		roleList = append(roleList, role)
	}

	return &types.SysRoleListResp{
		RoleList: roleList,
	}, nil
}

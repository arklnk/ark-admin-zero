package user

import (
	"ark-zero-admin/common/errorx"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserListLogic {
	return &GetSysUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserListLogic) GetSysUserList() (resp *types.SysUserListResp, err error) {
	users, err := l.svcCtx.SysUserModel.FindPage(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var user types.User
	var userList []types.User
	for _, v := range users {
		var roleIds []string
		err := copier.Copy(&user, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		err = json.Unmarshal([]byte(v.RoleIds), &roleIds)
		user.Roles = roleIds
		userList = append(userList, user)
	}

	return &types.SysUserListResp{UserList: userList}, nil
}

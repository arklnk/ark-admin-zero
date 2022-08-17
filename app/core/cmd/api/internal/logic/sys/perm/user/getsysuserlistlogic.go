package user

import (
	"ark-zero-admin/common/errorx"
	"context"
	"github.com/jinzhu/copier"
	"strings"

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
	users, err := l.svcCtx.SysUserModel.FindByPage(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var user types.User
	var userList []types.User
	for _, v := range users {
		err := copier.Copy(&user, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		user.Roles = strings.Split(v.Roles, ",")
		userList = append(userList, user)
	}

	return &types.SysUserListResp{UserList: userList}, nil
}

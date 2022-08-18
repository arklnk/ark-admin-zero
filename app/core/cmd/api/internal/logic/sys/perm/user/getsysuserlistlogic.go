package user

import (
	"context"
	"strconv"
	"strings"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
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

func (l *GetSysUserListLogic) GetSysUserList(req *types.SysUserListReq) (resp *types.SysUserListResp, err error) {
	var deptIds string
	for i, id := range req.DeptIds {
		if i == 0 {
			deptIds = strconv.FormatInt(id, 10)
			continue
		}
		deptIds = deptIds + "," + strconv.FormatInt(id, 10)
	}

	users, err := l.svcCtx.SysUserModel.FindByPage(l.ctx, req.Page, req.Limit, deptIds)
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

	total, err := l.svcCtx.SysUserModel.FindCountByDeptIds(l.ctx, deptIds)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var pagination types.Pagination
	pagination.Page = req.Page
	pagination.Size = req.Limit
	pagination.Total = total

	return &types.SysUserListResp{
		UserList:   userList,
		Pagination: pagination,
	}, nil
}

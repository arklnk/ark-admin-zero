package user

import (
	"ark-admin-zero/app/core/model"
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
	s := strconv.FormatInt(req.DeptId, 10)
	deptIds :=l.getDeptIds(s,req.DeptId)

	users, err := l.svcCtx.SysUserModel.FindByPage(l.ctx, req.Page, req.Limit, deptIds)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var user types.User
	var userProfession types.UserProfession
	var userJob types.UserJob
	var userDept types.UserDept
	userList := make([]types.User, 0)
	for _, v := range users {
		err := copier.Copy(&user, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}

		userProfession.Id = v.ProfessionId
		userProfession.Name = v.Profession

		userJob.Id = v.JobId
		userJob.Name = v.Job

		userDept.Id = v.DeptId
		userDept.Name = v.Dept

		var userRole types.UserRole
		var roles []types.UserRole
		var roleNameArr []string
		var roleIdArr []string
		roleNameArr = strings.Split(v.Roles, ",")
		roleIdArr = strings.Split(v.RoleIds, ",")
		for i, n := range roleNameArr {
			userRole.Name = n
			userRole.Id, _ = strconv.ParseInt(roleIdArr[i], 10, 64)
			roles = append(roles, userRole)
		}

		user.Profession = userProfession
		user.Job = userJob
		user.Dept = userDept
		user.Roles = roles

		userList = append(userList, user)
	}

	total, err := l.svcCtx.SysUserModel.FindCountByDeptIds(l.ctx, deptIds)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.UserListPagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysUserListResp{
		UserList:   userList,
		Pagination: pagination,
	}, nil
}

func (l *GetSysUserListLogic) getDeptIds(deptId string, id int64) string {
	deptList, err := l.svcCtx.SysDeptModel.FindSubDept(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return deptId
	}

	for _, v := range deptList {
		deptId = deptId + "," + strconv.FormatInt(v.Id, 10)
		deptId = l.getDeptIds(deptId, v.Id)
	}

	return deptId
}

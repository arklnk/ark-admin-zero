package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserRdpjInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserRdpjInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserRdpjInfoLogic {
	return &GetSysUserRdpjInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserRdpjInfoLogic) GetSysUserRdpjInfo(req *types.GetSysUserRdpjInfoReq) (resp *types.GetSysUserRdpjInfoResp, err error) {
	currentUserId := utils.GetUserId(l.ctx)
	return &types.GetSysUserRdpjInfoResp{
		Role:       l.roleList(currentUserId, req.UserId),
		Dept:       l.deptList(),
		Profession: l.professionList(),
		Job:        l.jobList(),
	}, nil
}

func (l *GetSysUserRdpjInfoLogic) roleList(currentUserId int64, editUserId int64) []types.RoleTree {
	var currentUserRoleIds []int64
	var sysRoleList []*model.SysRole
	sysRoleList, _ = l.svcCtx.SysRoleModel.FindAll(l.ctx)
	for _, role := range sysRoleList {
		currentUserRoleIds=append(currentUserRoleIds,role.Id)
	}
	
	var role types.RoleTree
	roleList := make([]types.RoleTree, 0)
	for _, v := range sysRoleList {
		_ = copier.Copy(&role, &v)
		roleList = append(roleList, role)
	}

	return roleList
}

func (l *GetSysUserRdpjInfoLogic) deptList() []types.DeptTree {
	sysDeptList, _ := l.svcCtx.SysDeptModel.FindEnable(l.ctx)
	var dept types.DeptTree
	deptList := make([]types.DeptTree, 0)
	for _, v := range sysDeptList {
		err := copier.Copy(&dept, &v)
		if err != nil {
			return nil
		}
		deptList = append(deptList, dept)
	}

	return deptList
}

func (l *GetSysUserRdpjInfoLogic) professionList() []types.Rdpj {
	sysProfessionList, _ := l.svcCtx.SysProfessionModel.FindEnable(l.ctx)
	var profession types.Rdpj
	professionList := make([]types.Rdpj, 0)
	for _, v := range sysProfessionList {
		err := copier.Copy(&profession, &v)
		if err != nil {
			return nil
		}
		professionList = append(professionList, profession)
	}

	return professionList
}

func (l *GetSysUserRdpjInfoLogic) jobList() []types.Rdpj {
	sysJobList, _ := l.svcCtx.SysJobModel.FindEnable(l.ctx)
	var job types.Rdpj
	jobList := make([]types.Rdpj, 0)
	for _, v := range sysJobList {
		err := copier.Copy(&job, &v)
		if err != nil {
			return nil
		}
		jobList = append(jobList, job)
	}

	return jobList
}

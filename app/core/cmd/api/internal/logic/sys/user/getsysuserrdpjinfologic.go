package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

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

func (l *GetSysUserRdpjInfoLogic) GetSysUserRdpjInfo() (resp *types.GetSysUserRdpjInfoResp, err error) {
	return &types.GetSysUserRdpjInfoResp{
		Role:       l.roleList(),
		Dept:       l.deptList(),
		Profession: l.professionList(),
		Job:        l.jobList(),
	}, nil
}

func (l *GetSysUserRdpjInfoLogic) roleList() []types.RdpjTree {
	sysRoleList, _ := l.svcCtx.SysRoleModel.FindEnable(l.ctx)
	var role types.RdpjTree
	roleList := make([]types.RdpjTree, 0)
	for _, v := range sysRoleList {
		err := copier.Copy(&role, &v)
		if err != nil {
			return nil
		}
		roleList = append(roleList, role)
	}

	return roleList
}

func (l *GetSysUserRdpjInfoLogic) deptList() []types.RdpjTree {
	sysDeptList, _ := l.svcCtx.SysDeptModel.FindEnable(l.ctx)
	var dept types.RdpjTree
	deptList := make([]types.RdpjTree, 0)
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

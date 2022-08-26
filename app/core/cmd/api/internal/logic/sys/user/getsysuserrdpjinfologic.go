package user

import (
	"context"
	"encoding/json"
	"strconv"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
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

func (l *GetSysUserRdpjInfoLogic) roleList(currentUserId uint64, editUserId uint64) []types.RdpjTree {
	currentUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, currentUserId)
	var currentUserRole []uint64
	err := json.Unmarshal([]byte(currentUser.RoleIds), &currentUserRole)
	if err != nil {
		return nil
	}

	editUser, _ := l.svcCtx.SysUserModel.FindOne(l.ctx, editUserId)
	var editUserRole []uint64
	err = json.Unmarshal([]byte(editUser.RoleIds), &editUserRole)
	if err != nil {
		return nil
	}

	var roleIds []uint64
	roleIds = append(roleIds, currentUserRole...)
	roleIds = append(roleIds, editUserRole...)
	var ids string
	for i, v := range roleIds {
		if i == 0 {
			ids = strconv.FormatUint(v, 10)
		}
		ids = ids + "," + strconv.FormatUint(v, 10)
	}

	sysRoleList, _ := l.svcCtx.SysRoleModel.FindByIds(l.ctx, ids)
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

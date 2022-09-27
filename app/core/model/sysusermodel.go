package model

import (
	"ark-admin-zero/config"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type SysUserDetail struct {
	Id           int64    `db:"id"`            // 编号
	Account      string    `db:"account"`       // 账号
	Username     string    `db:"username"`      // 姓名
	Nickname     string    `db:"nickname"`      // 昵称
	Avatar       string    `db:"avatar"`        // 头像
	Gender       int64    `db:"gender"`        // 0=保密 1=女 2=男
	Profession   string    `db:"profession"`    // 职称
	ProfessionId int64    `db:"profession_id"` // 职称id
	Job          string    `db:"job"`           // 岗位
	JobId        int64    `db:"job_id"`        // 岗位id
	Dept         string    `db:"dept"`          // 部门
	DeptId       int64    `db:"dept_id"`       // 部门id
	Roles        string    `db:"roles"`         // 角色集
	RoleIds      string    `db:"role_ids"`      // 角色集id
	Email        string    `db:"email"`         // 邮件
	Mobile       string    `db:"mobile"`        // 手机号
	Remark       string    `db:"remark"`        // 备注
	OrderNum     int64    `db:"order_num"`     // 排序值
	Status       int64    `db:"status"`        // 0=禁用 1=开启
	CreateTime   time.Time `db:"create_time"`   // 创建时间
	UpdateTime   time.Time `db:"update_time"`   // 更新时间
}

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindPage(ctx context.Context, page int64, limit int64, deptIds string) ([]*SysUserDetail, error)
		FindCountByCondition(ctx context.Context, condition string, value int64) (int64, error)
		FindCountByDeptIds(ctx context.Context, deptIds string) (int64, error)
		FindCountByRoleId(ctx context.Context, roleId int64) (int64, error)
		FindCountByJobId(ctx context.Context, jobId int64) (int64, error)
		FindCountByProfessionId(ctx context.Context, professionId int64) (int64, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}

func (m *customSysUserModel) FindPage(ctx context.Context, page int64, limit int64, deptIds string) ([]*SysUserDetail, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT u.id,u.dept_id,u.job_id,u.profession_id,u.account,u.username,u.nickname,u.avatar,u.gender,IFNULL(p.name,'NULL') as profession,IFNULL(j.name,'NULL') as job,IFNULL(d.name,'NULL') as dept,IFNULL(GROUP_CONCAT(r.name),'NULL') as roles,IFNULL(GROUP_CONCAT(r.id),0) as role_ids,u.email,u.mobile,u.remark,u.order_num,u.status,u.create_time,u.update_time FROM (SELECT * FROM sys_user WHERE id!=%d AND dept_id IN(%s) ORDER BY order_num DESC LIMIT %d,%d) u LEFT JOIN sys_profession p ON u.profession_id=p.id LEFT JOIN sys_dept d ON u.dept_id=d.id LEFT JOIN sys_job j ON u.job_id=j.id LEFT JOIN sys_role r ON JSON_CONTAINS(u.role_ids,JSON_ARRAY(r.id)) GROUP BY u.id", config.SysSuperUserId, deptIds, offset, limit)
	var resp []*SysUserDetail
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysUserModel) FindCountByCondition(ctx context.Context, condition string, value int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=%d", m.table, condition,value)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysUserModel) FindCountByDeptIds(ctx context.Context, deptIds string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE id!=%d AND dept_id IN(%s)", m.table, config.SysSuperUserId, deptIds)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysUserModel) FindCountByRoleId(ctx context.Context, roleId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s u WHERE JSON_CONTAINS(u.role_ids,JSON_ARRAY(%d))", m.table, roleId)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysUserModel) FindCountByJobId(ctx context.Context, jobId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE job_id=%d", m.table, jobId)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysUserModel) FindCountByProfessionId(ctx context.Context, jobId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE profession_id=%d", m.table, jobId)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

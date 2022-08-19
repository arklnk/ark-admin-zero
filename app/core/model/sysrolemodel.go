package model

import (
	"ark-admin-zero/common/globalkey"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		FindSubRole(ctx context.Context, id int64) ([]*SysRole, error)
		FindAll(ctx context.Context) ([]*SysRole, error)
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn, c),
	}
}

func (m *customSysRoleModel) FindSubRole(ctx context.Context, id int64) ([]*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where `parent_id` = ?", sysRoleRows, m.table)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysRoleModel) FindAll(ctx context.Context) ([]*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where id!=%d", sysRoleRows, m.table, globalkey.SysSuperAdminRoleId)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

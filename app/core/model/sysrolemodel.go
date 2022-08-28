package model

import (
	"ark-admin-zero/config"
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
		FindAll(ctx context.Context) ([]*SysRole, error)
		FindEnable(ctx context.Context) ([]*SysRole, error)
		FindByIds(ctx context.Context, ids string) ([]*SysRole, error)
		FindSubRole(ctx context.Context, id uint64) ([]*SysRole, error)
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

func (m *customSysRoleModel) FindAll(ctx context.Context) ([]*SysRole, error) {
	query := fmt.Sprintf("SELECT %s FROM %s  WHERE id!=%d ORDER BY order_num DESC", sysRoleRows, m.table, config.SysProtectRoleId)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysRoleModel) FindEnable(ctx context.Context) ([]*SysRole, error) {
	query := fmt.Sprintf("SELECT %s FROM %s  WHERE id!=%d AND status=1 ORDER BY order_num DESC", sysRoleRows, m.table, config.SysProtectRoleId)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysRoleModel) FindByIds(ctx context.Context, ids string) ([]*SysRole, error) {
	query := fmt.Sprintf("SELECT %s FROM %s  WHERE id!=%d AND status=1 AND id IN(%s) ORDER BY order_num DESC", sysRoleRows, m.table, config.SysProtectRoleId, ids)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysRoleModel) FindSubRole(ctx context.Context, id uint64) ([]*SysRole, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `parent_id`=%d", sysRoleRows, m.table, id)
	var resp []*SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

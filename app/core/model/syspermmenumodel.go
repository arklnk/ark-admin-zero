package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPermMenuModel = (*customSysPermMenuModel)(nil)

type (
	// SysPermMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPermMenuModel.
	SysPermMenuModel interface {
		sysPermMenuModel
		FindByIds(ctx context.Context, ids string) ([]*SysPermMenu, error)
		FindCountByParentId(ctx context.Context, id int64) (int64, error)
		FindAll(ctx context.Context) ([]*SysPermMenu, error)
		FindSubPermMenu(ctx context.Context, id int64) ([]*SysPermMenu, error)
	}

	customSysPermMenuModel struct {
		*defaultSysPermMenuModel
	}
)

// NewSysPermMenuModel returns a model for the database table.
func NewSysPermMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysPermMenuModel {
	return &customSysPermMenuModel{
		defaultSysPermMenuModel: newSysPermMenuModel(conn, c),
	}
}

func (m *customSysPermMenuModel) FindByIds(ctx context.Context, ids string) ([]*SysPermMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` in(%s)", sysPermMenuRows, m.table, ids)
	var resp []*SysPermMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysPermMenuModel) FindCountByParentId(ctx context.Context, id int64) (int64, error) {
	query := fmt.Sprintf("select count(id) from %s where `parent_id`=%d", m.table, id)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysPermMenuModel) FindAll(ctx context.Context) ([]*SysPermMenu, error) {
	query := fmt.Sprintf("select %s from %s order by order_num desc", sysPermMenuRows, m.table)
	var resp []*SysPermMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysPermMenuModel) FindSubPermMenu(ctx context.Context, id int64) ([]*SysPermMenu, error) {
	query := fmt.Sprintf("select %s from %s where `parent_id` = ?", sysPermMenuRows, m.table)
	var resp []*SysPermMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysConfigModel = (*customSysConfigModel)(nil)

type (
	// SysConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysConfigModel.
	SysConfigModel interface {
		sysConfigModel
		FindList(ctx context.Context) ([]*SysConfig, error)
		FindPageByParentId(ctx context.Context, id int64, page int64, limit int64) ([]*SysConfig, error)
		FindCountByParentId(ctx context.Context, id int64) (int64, error)
	}

	customSysConfigModel struct {
		*defaultSysConfigModel
	}
)

// NewSysConfigModel returns a model for the database table.
func NewSysConfigModel(conn sqlx.SqlConn, c cache.CacheConf) SysConfigModel {
	return &customSysConfigModel{
		defaultSysConfigModel: newSysConfigModel(conn, c),
	}
}

func (m *customSysConfigModel) FindList(ctx context.Context) ([]*SysConfig, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE parent_id=0 ORDER BY order_num DESC", sysConfigRows, m.table)
	var resp []*SysConfig
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysConfigModel) FindPageByParentId(ctx context.Context, id int64, page int64, limit int64) ([]*SysConfig, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT %s FROM %s WHERE parent_id=? ORDER BY order_num DESC LIMIT %d,%d", sysConfigRows, m.table, offset, limit)
	var resp []*SysConfig
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysConfigModel) FindCountByParentId(ctx context.Context, id int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE parent_id=?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

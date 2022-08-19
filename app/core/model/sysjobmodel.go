package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysJobModel = (*customSysJobModel)(nil)

type (
	// SysJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysJobModel.
	SysJobModel interface {
		sysJobModel
		FindAll(ctx context.Context) ([]*SysJob, error)
		FindByPage(ctx context.Context, page int64, limit int64) ([]*SysJob, error)
		FindCount(ctx context.Context) (int64, error)
	}

	customSysJobModel struct {
		*defaultSysJobModel
	}
)

// NewSysJobModel returns a model for the database table.
func NewSysJobModel(conn sqlx.SqlConn, c cache.CacheConf) SysJobModel {
	return &customSysJobModel{
		defaultSysJobModel: newSysJobModel(conn, c),
	}
}

func (m *customSysJobModel) FindAll(ctx context.Context) ([]*SysJob, error) {
	query := fmt.Sprintf("select %s from %s order by order_num desc", sysJobRows, m.table)
	var resp []*SysJob
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *customSysJobModel) FindByPage(ctx context.Context, page int64, limit int64) ([]*SysJob, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("select %s from %s order by order_num desc limit %d,%d", sysJobRows, m.table, offset, limit)
	var resp []*SysJob
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysJobModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(id) from %s", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

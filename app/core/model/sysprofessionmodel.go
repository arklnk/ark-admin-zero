package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysProfessionModel = (*customSysProfessionModel)(nil)

type (
	// SysProfessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysProfessionModel.
	SysProfessionModel interface {
		sysProfessionModel
		FindAll(ctx context.Context) ([]*SysProfession, error)
		FindEnable(ctx context.Context) ([]*SysProfession, error)
		FindCount(ctx context.Context) (uint64, error)
		FindPage(ctx context.Context, page uint64, limit uint64) ([]*SysProfession, error)
	}

	customSysProfessionModel struct {
		*defaultSysProfessionModel
	}
)

// NewSysProfessionModel returns a model for the database table.
func NewSysProfessionModel(conn sqlx.SqlConn, c cache.CacheConf) SysProfessionModel {
	return &customSysProfessionModel{
		defaultSysProfessionModel: newSysProfessionModel(conn, c),
	}
}

func (m *customSysProfessionModel) FindAll(ctx context.Context) ([]*SysProfession, error) {
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY order_num DESC", sysProfessionRows, m.table)
	var resp []*SysProfession
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysProfessionModel) FindEnable(ctx context.Context) ([]*SysProfession, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE status=1 ORDER BY order_num DESC", sysProfessionRows, m.table)
	var resp []*SysProfession
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysProfessionModel) FindPage(ctx context.Context, page uint64, limit uint64) ([]*SysProfession, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY order_num DESC LIMIT %d,%d", sysProfessionRows, m.table, offset, limit)
	var resp []*SysProfession
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysProfessionModel) FindCount(ctx context.Context) (uint64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", m.table)
	var resp uint64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

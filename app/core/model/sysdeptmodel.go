package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDeptModel = (*customSysDeptModel)(nil)

type (
	// SysDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDeptModel.
	SysDeptModel interface {
		sysDeptModel
		FindAll(ctx context.Context) ([]*SysDept, error)
		FindSubDept(ctx context.Context, id int64) ([]*SysDept, error)
	}

	customSysDeptModel struct {
		*defaultSysDeptModel
	}
)

// NewSysDeptModel returns a model for the database table.
func NewSysDeptModel(conn sqlx.SqlConn, c cache.CacheConf) SysDeptModel {
	return &customSysDeptModel{
		defaultSysDeptModel: newSysDeptModel(conn, c),
	}
}

func (m *customSysDeptModel) FindAll(ctx context.Context) ([]*SysDept, error) {
	query := fmt.Sprintf("select %s from %s", sysDeptRows, m.table)
	var resp []*SysDept
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysDeptModel) FindSubDept(ctx context.Context, id int64) ([]*SysDept, error) {
	query := fmt.Sprintf("select %s from %s where `parent_id` = ?", sysDeptRows, m.table)
	var resp []*SysDept
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

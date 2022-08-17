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
	query := fmt.Sprintf("select %s from %s", sysProfessionRows, m.table)
	var resp []*SysProfession
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

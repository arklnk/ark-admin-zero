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
	query := fmt.Sprintf("select %s from %s", sysJobRows, m.table)
	var resp []*SysJob
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

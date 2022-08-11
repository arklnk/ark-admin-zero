package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPermMenuModel = (*customSysPermMenuModel)(nil)

type (
	// SysPermMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPermMenuModel.
	SysPermMenuModel interface {
		sysPermMenuModel
		FindAll(ctx context.Context, id int64) ([]*SysPermMenu, error)
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

func (m *customSysPermMenuModel) FindAll(ctx context.Context, id int64) ([]*SysPermMenu, error) {
	return nil, nil
}

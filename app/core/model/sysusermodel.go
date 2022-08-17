package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindByCondition(ctx context.Context, condition string, value int64) ([]*SysUser, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}

func (m *customSysUserModel) FindByCondition(ctx context.Context, condition string, value int64) ([]*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where %s=?", sysUserRows, m.table,condition)
	var resp []*SysUser
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, value)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

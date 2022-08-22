package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDictionaryModel = (*customSysDictionaryModel)(nil)

type (
	// SysDictionaryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionaryModel.
	SysDictionaryModel interface {
		sysDictionaryModel
		FindDictionarySet(ctx context.Context) ([]*SysDictionary, error)
		FindPageByParentId(ctx context.Context, id int64, page int64, limit int64) ([]*SysDictionary, error)
		FindCountByParentId(ctx context.Context, id int64) (int64, error)
	}

	customSysDictionaryModel struct {
		*defaultSysDictionaryModel
	}
)

// NewSysDictionaryModel returns a model for the database table.
func NewSysDictionaryModel(conn sqlx.SqlConn, c cache.CacheConf) SysDictionaryModel {
	return &customSysDictionaryModel{
		defaultSysDictionaryModel: newSysDictionaryModel(conn, c),
	}
}

func (m *customSysDictionaryModel) FindDictionarySet(ctx context.Context) ([]*SysDictionary, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE parent_id=0 ORDER BY order_num DESC", sysConfigRows, m.table)
	var resp []*SysDictionary
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysDictionaryModel) FindPageByParentId(ctx context.Context, id int64, page int64, limit int64) ([]*SysDictionary, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT %s FROM %s WHERE parent_id=? ORDER BY order_num DESC LIMIT %d,%d", sysConfigRows, m.table, offset, limit)
	var resp []*SysDictionary
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysDictionaryModel) FindCountByParentId(ctx context.Context, id int64) (int64, error) {
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

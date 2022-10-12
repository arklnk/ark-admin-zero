package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ SysLogModel = (*customSysLogModel)(nil)

type SysLoginLog struct {
	Id         int64     `db:"id"`          // 编号
	Account    string    `db:"account"`     // 操作账号
	Ip         string    `db:"ip"`          // ip
	Uri        string    `db:"uri"`         // 请求路径
	Status     int64     `db:"status"`      // 0=失败 1=成功
	CreateTime time.Time `db:"create_time"` // 创建时间
}

type (
	// SysLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLogModel.
	SysLogModel interface {
		sysLogModel
		FindPage(ctx context.Context, t int64, page int64, limit int64) ([]*SysLoginLog, error)
		FindCount(ctx context.Context, t int64) (int64, error)
	}

	customSysLogModel struct {
		*defaultSysLogModel
	}
)

// NewSysLogModel returns a model for the database table.
func NewSysLogModel(conn sqlx.SqlConn, c cache.CacheConf) SysLogModel {
	return &customSysLogModel{
		defaultSysLogModel: newSysLogModel(conn, c),
	}
}

func (m *customSysLogModel) FindPage(ctx context.Context, t int64, page int64, limit int64) ([]*SysLoginLog, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT l.id,IFNULL(u.account,'NULL') as account,l.ip,l.uri,l.status,l.create_time FROM (SELECT * FROM sys_log WHERE type=%d ORDER BY id DESC LIMIT %d,%d) l LEFT JOIN sys_user u ON l.user_id=u.id", t, offset, limit)
	var resp []*SysLoginLog
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysLogModel) FindCount(ctx context.Context, t int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE type=%d", m.table, t)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

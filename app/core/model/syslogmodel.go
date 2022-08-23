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
	Id         uint64    `db:"id"`          // 编号
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
		FindPageByType(ctx context.Context, t int64, page int64, limit int64) ([]*SysLoginLog, error)
		FindCountByType(ctx context.Context, t int64) (int64, error)
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

func (m *customSysLogModel) FindPageByType(ctx context.Context, t int64, page int64, limit int64) ([]*SysLoginLog, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT l.id,u.account,l.ip,l.uri,l.status,l.create_time FROM (SELECT * FROM sys_log WHERE type=? ORDER BY id DESC LIMIT ?,?) l LEFT JOIN sys_user u ON l.user_id=u.id")
	var resp []*SysLoginLog
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, t, offset, limit)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysLogModel) FindCountByType(ctx context.Context, t int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE type=?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, t)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

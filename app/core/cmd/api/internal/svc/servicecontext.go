package svc

import (
	"ark-zero-admin/app/core/cmd/api/internal/config"
	"ark-zero-admin/app/core/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	Redis            *redis.Redis
	SysUserModel     model.SysUserModel
	SysPermMenuModel model.SysPermMenuModel
	SysRoleModel     model.SysRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:           c,
		Redis:            redisClient,
		SysUserModel:     model.NewSysUserModel(mysqlConn, c.Cache),
		SysPermMenuModel: model.NewSysPermMenuModel(mysqlConn, c.Cache),
		SysRoleModel:     model.NewSysRoleModel(mysqlConn, c.Cache),
	}
}

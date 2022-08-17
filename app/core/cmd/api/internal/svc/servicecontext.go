package svc

import (
	"ark-zero-admin/app/core/cmd/api/internal/config"
	"ark-zero-admin/app/core/cmd/api/internal/middleware"
	"ark-zero-admin/app/core/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	Redis              *redis.Redis
	PermMenuAuth       rest.Middleware
	SysUserModel       model.SysUserModel
	SysPermMenuModel   model.SysPermMenuModel
	SysRoleModel       model.SysRoleModel
	SysDeptModel       model.SysDeptModel
	SysJobModel        model.SysJobModel
	SysProfessionModel model.SysProfessionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:             c,
		Redis:              redisClient,
		PermMenuAuth:       middleware.NewPermMenuAuthMiddleware(c.JwtAuth.AccessSecret, redisClient).Handle,
		SysUserModel:       model.NewSysUserModel(mysqlConn, c.Cache),
		SysPermMenuModel:   model.NewSysPermMenuModel(mysqlConn, c.Cache),
		SysRoleModel:       model.NewSysRoleModel(mysqlConn, c.Cache),
		SysDeptModel:       model.NewSysDeptModel(mysqlConn, c.Cache),
		SysJobModel:        model.NewSysJobModel(mysqlConn, c.Cache),
		SysProfessionModel: model.NewSysProfessionModel(mysqlConn, c.Cache),
	}
}

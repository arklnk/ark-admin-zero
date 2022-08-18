package main

import (
	"flag"
	"fmt"
	"net/http"

	"ark-admin-zero/app/core/cmd/api/internal/config"
	"ark-admin-zero/app/core/cmd/api/internal/handler"
	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	if c.Mode == "dev" {
		logx.DisableStat()
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

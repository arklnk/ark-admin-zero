package main

import (
	"flag"
	"fmt"
	"net/http"

    "ark-zero-admin/common/errorx"
	{{.importPackages}}
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

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

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

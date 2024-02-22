package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"looklook/app/shop/cmd/api/internal/config"
	"looklook/app/shop/cmd/api/internal/handler"
	"looklook/app/shop/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/shop.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//设置允许跨域访问
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

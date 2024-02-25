package main

import (
	"flag"
	"fmt"
	"looklook/app/comment/cmd/api/internal/config"
	"looklook/app/comment/cmd/api/internal/handler"
	"looklook/app/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//设置允许跨域访问
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//禁止性能监控日志 建议打卡
	//logx.DisableStat()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

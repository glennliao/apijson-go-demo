package main

import (
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/framework"
	"github.com/glennliao/apijson-go/framework/handler"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	framework.Init()

	config.AccessVerify = false // 全局配置验证权限开关

	s := g.Server()

	s.Group("/", handler.Bind)

	s.Run()
}

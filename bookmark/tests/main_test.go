package tests

import (
	"apijson-go-demo/bookmark/app"
	"context"
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/framework"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/os/gfile"
)

func init() {

	gfile.Remove("./bookmark.sqlite3")

	ctx := context.Background()
	app.Init(ctx)

	framework.Init()

	config.NoAccessVerify = true
	config.AccessConditionFunc = app.AccessCondition
	config.DefaultRoleFunc = app.Role
	config.AddRole(app.RoleGroupAdmin)
	config.AddRole(app.RoleGroupUser)

	initData()
}

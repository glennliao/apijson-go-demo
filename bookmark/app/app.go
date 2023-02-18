package app

import (
	"context"
	_ "embed"
	"github.com/glennliao/apijson-go/action"
	"github.com/glennliao/apijson-go/framework"
	"github.com/glennliao/apijson-go/framework/database"
	"github.com/glennliao/apijson-go/model"
	"github.com/glennliao/table-sync/tablesync"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
	"net/http"
)

const (
	TableUser          = "user"
	TableBookmark      = "bookmark"
	TableBookmarkCate  = "bookmark_cate"
	TableGroup         = "group"
	TableGroupUser     = "group_user"
	TableGroupBookmark = "group_bookmark"
)

//go:embed _access.json
var _access string

//go:embed _reuqest.json
var _request string

func Init(ctx context.Context) {
	initDB(ctx)
}

func initDB(ctx context.Context) {
	db := g.DB()

	tables, err := db.Tables(ctx)
	if err != nil {
		panic(err)
	}

	hasTable := lo.Contains(tables, "bookmark")

	syncer := tablesync.Syncer{Tables: Tables()}
	err = syncer.Sync(ctx, db)
	if err != nil {
		panic(err)
	}

	// init access
	if !hasTable {

		var accessList []database.Access
		err := gconv.Scan(_access, &accessList)
		if err != nil {
			panic(err)
		}
		db.Model("_access").Ctx(ctx).FieldsEx("id").Insert(accessList)

		var requestList []database.Request
		err = gconv.Scan(_request, &requestList)
		if err != nil {
			panic(err)
		}
		db.Model("_request").Ctx(ctx).FieldsEx("id").Insert(requestList)
	}

	// init data

	if !hasTable {

		framework.Init()

		a := action.New(ctx, http.MethodPost, model.Map{
			"tag": "User",
			"User": model.Map{
				"username": "admin",
			},
		})
		a.NoAccessVerify = true
		_, err := a.Result()

		if err != nil {
			g.Log().Fatal(ctx, err)
		}
	}
}

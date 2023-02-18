package tests

import (
	"apijson-go-demo/bookmark/app"
	"context"
	"github.com/glennliao/apijson-go/framework/handler"
	"github.com/glennliao/apijson-go/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

const (
	UsernameUser1 = "admin"
	UsernameUser2 = "user2"
	UsernameUser3 = "user3"
)

var userIdMap = map[string]string{
	UsernameUser1: "1",
}

func initData() {
	ctx := iamUser(UsernameUser1)
	addUser(ctx, UsernameUser2)
	addUser(ctx, UsernameUser3)
	addSomeData(ctx)
}

func iamUser(username string) context.Context {

	if _, exists := userIdMap[username]; !exists {
		one, err := g.DB().Ctx(context.Background()).Model("user").Where(g.Map{
			"username": username,
		}).One()
		AssertErrIsNil(nil, err)
		if one.IsEmpty() {
			panic("user not found:" + username)
		}
		userIdMap[username] = one.Map()["user_id"].(string)

	}

	ctx := context.WithValue(context.Background(), app.UserIdKey, &app.CurrentUser{
		UserId: userIdMap[username],
	})
	return ctx
}

func addUser(ctx context.Context, username string) {
	_, err := handler.Post(ctx, model.Map{
		"User": g.Map{
			"username": username,
		},
		"tag": "User",
	})

	if err != nil {
		g.Log().Error(ctx, err)
	}

	gtest.Assert(err, nil)
}

func addSomeData(ctx context.Context) {
	ctx = iamUser(UsernameUser2)
	groupId := userIdMap[UsernameUser2]
	cateId := bookmarkCateAdd(nil, ctx, groupId, "testCate2", "root")
	subCateId := bookmarkCateAdd(nil, ctx, groupId, "testSubCate2", cateId)
	bookmarkAdd(nil, ctx, groupId, subCateId, "https://github.com", "github( by 2)")

	ctx = iamUser(UsernameUser3)
	groupId = userIdMap[UsernameUser3]
	cateId = bookmarkCateAdd(nil, ctx, groupId, "testCate3", "root")
	subCateId = bookmarkCateAdd(nil, ctx, groupId, "testSubCate3", cateId)
	bookmarkAdd(nil, ctx, groupId, subCateId, "https://github.com", "github( by 3)")
}

func AssertErrIsNil(t *gtest.T, err error) {
	if t != nil {
		t.AssertNil(err)
	} else {
		if err != nil {
			panic(err)
		}
	}
}

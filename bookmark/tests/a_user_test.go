package tests

import (
	"context"
	"github.com/glennliao/apijson-go/framework/handler"
	"github.com/glennliao/apijson-go/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

// 普通用户
func TestUser(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		// add cate,bookmark
		ctx := iamUser(UsernameUser3)
		groupId := userIdMap[UsernameUser3]
		cateId := bookmarkCateAdd(t, ctx, groupId, "testCate3-2", "root")
		subCateId := bookmarkCateAdd(t, ctx, groupId, "testSubCate3-2_1", cateId)
		bookmarkAdd(t, ctx, groupId, subCateId, "https://www.gitee.com", "gitee")
	})

	gtest.C(t, func(t *gtest.T) {
		// 查看所在组的书签
		ctx := iamUser(UsernameUser3)

		ret, err := handler.Get(ctx, model.Map{
			"BookmarkCate[]": model.Map{
				"@column": "title,parentId,cateId,groupId,createdBy",
			},
		})
		t.AssertNil(err)
		g.Dump(ret)
		ret, err = handler.Get(ctx, model.Map{
			"Bookmark[]": model.Map{
				"@column": "title,url,createdBy",
				"cateId":  ret["BookmarkCate[]"].([]model.Map)[1]["cateId"].(string),
			},
		})
		t.AssertNil(err)
		g.Dump(ret)
	})

}

func bookmarkCateAdd(t *gtest.T, ctx context.Context, groupId string, cateTitle string, parentId string) (cateId string) {
	ret, err := handler.Post(ctx, model.Map{
		"BookmarkCate": g.Map{
			"parentId": parentId,
			"title":    cateTitle,
			"groupId":  groupId,
		},
		"tag": "BookmarkCate",
	})
	AssertErrIsNil(t, err)
	cateId = ret["BookmarkCate"].(model.Map)["cateId"].(string)
	return
}

func bookmarkAdd(t *gtest.T, ctx context.Context, groupId, cateId, url, title string) (bookmarkId string) {
	_, err := handler.Post(ctx, model.Map{
		"Bookmark": g.Map{
			"url":    url,
			"title":  title,
			"cateId": cateId,
		},
		"GroupBookmark": g.Map{
			"cateId":  cateId,
			"groupId": groupId,
			"bmId@":   "Bookmark/bmId",
		},
		"tag": "Bookmark",
	})

	AssertErrIsNil(t, err)
	return ""
}

package app

import (
	"context"
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/consts"
	"github.com/glennliao/apijson-go/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/samber/lo"
	"net/http"
)

const UserIdKey = "userId"

type CurrentUser struct {
	UserId string
}

const RoleGroupUser = "GroupUser"
const RoleGroupAdmin = "GroupAdmin"

func Role(ctx context.Context, req config.RoleReq) (string, error) {
	_, ok := ctx.Value(UserIdKey).(*CurrentUser)

	if !ok {
		return consts.UNKNOWN, nil //未登录
	}

	if req.NodeRole == "" {

		switch req.Table {
		case TableUser:
			return consts.OWNER, nil
		}

	} else {

		switch req.Table {
		case TableUser:
			if req.NodeRole == consts.LOGIN {
				return consts.OWNER, nil
			}

		case TableBookmark:

			if req.NodeRole == consts.LOGIN {
				req.NodeRole = consts.OWNER
			}

			if lo.Contains([]string{consts.OWNER, RoleGroupUser}, req.NodeRole) {
				return req.NodeRole, nil
			}

			return consts.DENY, nil // 非拥有的角色

		default:
			return req.NodeRole, nil
		}
	}

	return consts.LOGIN, nil

}

func AccessCondition(ctx context.Context, req config.AccessConditionReq) (where g.Map, err error) {

	where = g.Map{}
	rawSql := model.Map{}
	user, ok := ctx.Value(UserIdKey).(*CurrentUser)

	if !ok {
		return nil, nil
	}

	switch req.Table {
	case TableUser:
		if req.NodeRole == consts.OWNER {
			return g.Map{
				"user_id": user.UserId,
			}, nil
		}
	case TableBookmarkCate:
		return g.Map{
			consts.Raw: model.Map{
				"group_id in (select group_id from group_user where user_id = ?)": user.UserId,
			},
		}, nil
	case TableBookmark:
		if req.Method == http.MethodGet {
			if v, exists := req.NodeReq["cateId"]; exists {
				delete(req.NodeReq, "cateId")
				// 性能? 拓展?
				rawSql["bm_id in (select bm_id from group_bookmark where cate_id = ? and group_id in (select group_id from group_user where user_id = ? ))"] = []string{v.(string), user.UserId}
			} else {
				rawSql["bm_id in (select bm_id from group_bookmark where  group_id in (select group_id from group_user where user_id = ? ))"] = []string{user.UserId}
			}
		}

	case "t_user":
		if req.NodeRole == consts.OWNER {
			return g.Map{
				"user_id": user.UserId,

				consts.Raw: g.Map{
					"id > 0": "",
				},
				// g.Map{
				//    "uid <=" : 1000,
				//    "age >=" : 18,
				//    "x in (?)":[]string{"1","2","3"}
				//}
			}, nil
		}
	case "t_todo":
		if req.NodeRole == consts.OWNER {
			return g.Map{"user_id": user.UserId}, nil
		}
		if req.NodeRole == RoleGroupUser {
			return g.Map{"partner": user.UserId}, nil
		}
	}

	where[consts.Raw] = rawSql

	return where, nil
}

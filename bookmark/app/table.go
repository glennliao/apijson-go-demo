package app

import (
	"github.com/glennliao/apijson-go/framework/database"
	"github.com/glennliao/table-sync/tablesync"
	"time"
)

type Bookmark struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"书签"`
	Id                  uint64 `ddl:"primaryKey"`
	BmId                string `ddl:"size:32;comment:书签id"`
	Title               string `ddl:"size:128;comment:标题"`
	Url                 string `ddl:"size:512;comment:书签地址"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

type BookmarkCate struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"书签目录"`
	Id                  uint64 `ddl:"primaryKey"`
	CateId              string `ddl:"size:32;comment:目录id"`
	Title               string `ddl:"size:128;comment:群组名"`
	ParentId            string `ddl:"size:32;comment:父级目录id"`
	GroupId             string `ddl:"size:32;comment:群组id"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

type GroupBookmark struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"书签目录"`
	Id                  uint64 `ddl:"primaryKey"`
	CateId              string `ddl:"size:32;comment:目录id"`
	BmId                string `ddl:"size:32;comment:书签id"`
	GroupId             string `ddl:"size:32;comment:群组id"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

type Group struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"群组"`
	Id                  uint64 `ddl:"primaryKey"`
	GroupId             string `ddl:"size:32;comment:群组id"`
	Title               string `ddl:"size:128;comment:群组名"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

type GroupUser struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"群组用户"`
	Id                  uint64 `ddl:"primaryKey"`
	GroupId             string `ddl:"size:32;comment:群组id"`
	UserId              string `ddl:"size:32;comment:用户id"`
	IsAdmin             uint8  `ddl:"comment:是否管理员"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

type User struct {
	tablesync.TableMeta `charset:"utf8mb4" comment:"用户"`
	Id                  uint64 `ddl:"primaryKey"`
	UserId              string `ddl:"size:32;comment:用户id"`
	Username            string `ddl:"size:512;comment:用户名"`
	CreatedAt           *time.Time
	CreatedBy           string `ddl:"size:32"`
	UpdatedAt           *time.Time
	UpdatedBy           string `ddl:"size:32"`
	DeletedAt           *time.Time
}

func Tables() []tablesync.Table {
	return []tablesync.Table{
		// app
		Bookmark{},
		BookmarkCate{},
		Group{},
		GroupUser{},
		GroupBookmark{},
		User{},
		// apijson
		database.Access{},
		database.Request{},
	}
}

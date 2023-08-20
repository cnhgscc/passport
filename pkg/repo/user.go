package repo

import (
	dbp "github.com/cnhgscc/passport/pkg/dbplugin"
)

// User 用户信息
type User struct {
	dbp.M

	NickName string `gorm:"size:32"`
	Age      int
	Avatar   string
	Gender   int
}

func (u *User) TableName() string {
	return "user"
}

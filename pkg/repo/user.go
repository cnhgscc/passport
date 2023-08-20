package repo

import (
	dbp "github.com/cnhgscc/passport/pkg/dbplugin"
)

// Login 登录信息
type Login struct {
	dbp.M

	// 注册数据
	Method   string
	Username string
	Password string
	RegTS    int64
	RegIP    string
	RegMAC   string

	// 登录数据
	LastLoginTS  int64
	LastLoginIP  string
	LastLoginMAC string
}

// User 用户信息
type User struct {
	dbp.M

	NickName string
	Age      int
	Avatar   string
	Gender   int
}

func (u *User) TableName() string {
	return "user"
}

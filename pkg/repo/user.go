package repo

import (
	dbp "github.com/cnhgscc/passport/pkg/dbplugin"
)

type Passport struct {
	dbp.M

	Username string
	Password string
}

type User struct {
	dbp.M

	NickName string
	Age      int
}

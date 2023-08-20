package repo

import (
	"fmt"

	dbp "github.com/cnhgscc/passport/pkg/dbplugin"
)

// Login 登录信息
type Login struct {
	dbp.M
	// 注册数据
	Method   string `gorm:"size:32"`
	Username string `gorm:"suze:128"`
	Password string `gorm:"size:32"`
	Regts    int64
	Regip    string `gorm:"size:32"`
	Regmac   string `grom:"size:64"`
	// 登录数据
	LastLogints  int64
	LastLoginip  string `gorm:"32"`
	LastLoginmac string `gorm:"64"`
}

func (l *Login) TableName() string {
	return "login"
}

func (l *Login) Auth() error {

	db := l.DB.Where("method=? AND username=?", l.Method, l.Username)
	if l.Password != "" {
		db = db.Where("password=?", l.Password)
	}
	l.Err = db.Limit(1).Find(l).Error
	if l.Err != nil {
		return nil
	}

	if l.ID == 0 && l.Password != "" {
		return fmt.Errorf("password err")
	}

	return nil
}

func NewLoginAccount(method, username string, opts ...LoginOption) *Login {
	login := &Login{Method: method, Username: username}
	for _, opt := range opts {
		opt(login)
	}
	return login
}

type LoginOption func(login *Login)

func WithPassword(password string) LoginOption {
	return func(login *Login) {
		// TODO: hasher
		login.Password = password
	}
}

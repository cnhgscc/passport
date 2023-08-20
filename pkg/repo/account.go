package repo

import (
	"fmt"
	"time"

	dbp "github.com/cnhgscc/passport/pkg/dbplugin"
)

// Account 登录信息
type Account struct {
	dbp.M
	Method       string `gorm:"size:32"`
	Username     string `gorm:"suze:128"`
	Password     string `gorm:"size:32"`
	Regts        int64
	Regip        string `gorm:"size:32"`
	Regmac       string `grom:"size:64"`
	LastLogints  int64
	LastLoginip  string `gorm:"32"`
	LastLoginmac string `gorm:"64"`
}

func (l *Account) TableName() string {
	return "account"
}

func (l *Account) Auth(db *dbp.DB) error {
	db = db.Where("method=? AND username=?", l.Method, l.Username)
	if l.Password != "" {
		db = db.Where("password=?", l.Password)
	}
	err := db.Limit(1).Find(l).Error
	if err != nil {
		return fmt.Errorf("system err")
	}
	if l.Password != "" && l.ID != 0 {
		return fmt.Errorf("password err")
	}
	if l.ID != 0 {
		return nil
	}
	l.Regts = l.LastLogints
	l.Regip = l.LastLoginip
	l.Regmac = l.LastLoginmac
	db.Save(l)
	return nil
}

func NewLoginAccount(method, username, ip, mac string, opts ...LoginOption) *Account {
	login := &Account{Method: method, Username: username, LastLoginip: ip, LastLoginmac: mac}
	login.LastLogints = time.Now().UnixMilli()
	for _, opt := range opts {
		opt(login)
	}
	return login
}

type LoginOption func(ua *Account)

func WithPassword(password string) LoginOption {
	return func(login *Account) {
		// TODO: hasher
		login.Password = password
	}
}

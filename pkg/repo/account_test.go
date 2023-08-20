package repo

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cnhgscc/passport/pkg/gomysql"
)

func init() {

	err := gomysql.Init(&gomysql.Config{
		Scope: "passport",
		User:  "root",
		Pass:  "123456",
		DB:    "passport",
		Addr:  "127.0.0.1",
		Max:   10,
	})
	if err != nil {
		fmt.Println(err)
	}

}

func TestMySQL(t *testing.T) {

	db := gomysql.S("passport")
	db.AutoMigrate(&Account{})

	fmt.Println(db)
	l := NewLoginAccount("google", "root1", "127.0.0.1", "xdas123123")

	err := l.Auth(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	marshal, err := json.Marshal(l)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	fmt.Println(l.ID)

}

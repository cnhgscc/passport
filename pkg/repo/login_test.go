package repo

import (
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
	fmt.Println(db)
	db.AutoMigrate(&Login{})

}

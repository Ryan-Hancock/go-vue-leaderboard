package storage

import (
	"fmt"

	"github.com/Ryan-Hancock/go-vue-leaderboard/context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	ctx := context.GetContext()
	var err error
	fmt.Println(ctx.Config.DBString)
	db, err = gorm.Open("mysql", ctx.Config.DBString)

	if err != nil {
		panic(err)
	}
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

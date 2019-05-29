package models

import (
	"github.com/Ryan-Hancock/go-vue-leaderboard/storage"
)

var MigrationMap = []interface{}{
	&Score{},
}

func Init() {
	db := storage.GetDB()

	for _, i := range MigrationMap {
		db.AutoMigrate(i)
	}
}

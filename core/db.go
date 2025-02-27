package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbFile string = "chat.db"

var DB *gorm.DB

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Message{},
		&User{},
	)
}

func InitDB(persistent bool) error {
	url := "file::memory:?cache=shared"
	if persistent {
		url = dbFile
	}

	gdb, err := gorm.Open(sqlite.Open(url))
	if err != nil {
		return err
	}
	DB = gdb

	return AutoMigrate(DB)
}
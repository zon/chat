package core

import (
	"github.com/zon/gonf"
)

const defaultDatabase = "chat"
const sqliteFile = defaultDatabase + ".db"

func InitDB() error {
	return gonf.InitDB(defaultDatabase, sqliteFile)
}

func AutoMigrate() error {
	err := gonf.AutoMigrate()
	if err != nil {
		return err
	}
	return gonf.DB.AutoMigrate(
		&Message{},
	)
}

package core

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
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

func InitDB() error {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	database := os.Getenv("PGDATABASE")

	var dialector gorm.Dialector
	if host != "" {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s",
			host, port, user, password, database,
		)
		dialector = postgres.Open(dsn)
	} else {
		dialector = sqlite.Open(dbFile)
	}

	gdb, err := gorm.Open(dialector, &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	DB = gdb

	return AutoMigrate(DB)
}

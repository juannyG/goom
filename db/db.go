package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB) {
	dsn := os.Getenv("GOOM_DSN")
	if dsn == "" {
		panic("No DSN defined - cannot connect to the database")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("There was an error connecting to the database - %s", err))
	}
	return db
}

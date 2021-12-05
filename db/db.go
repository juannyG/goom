package db

import (
	"database/sql"
)


var DB *sql.DB

func Connect() error {
	return nil
}

func GetDSN() (string, error) {
	return "", nil
}

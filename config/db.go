package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", "./driversapi.db")
	return DB, err
}

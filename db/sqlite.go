package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitSqliteConnection() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "./homelobby.db")
	if err != nil {
		panic(err)
	} else {
		println("SQLite connected.")
	}

	return db
}

package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func conn() (*sql.DB, error) {
	var err error
	conn, err := sql.Open("sqlite3", "Tasks.db")
	if err != nil {
		log.Fatal("DB connection failed: " + err.Error())
		return nil, err
	}

	return conn, nil
}

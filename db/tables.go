package db

import "log"

func CreateTasks() error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()
	sql := `
CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY,
	title TEXT,
	description TEXT,
	due_date DATE,
	status TEXT
	);`

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return err
}

func CreateUsers() error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	sql := `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	email TEXT UNIQUE,
	hash TEXT
	);`

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

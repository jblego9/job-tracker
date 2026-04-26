package repository

import (
	"database/sql"
	"os"
)

var db *sql.DB

func Init() error {
	if err := open(); err != nil {
		return err
	}

	if err := setup(); err != nil {
		return err
	}

	return nil
}

func dbPath() string {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "../data/dev.db"
	}

	return dbPath
}

func open() error {
	new_db, err := sql.Open("sqlite", dbPath()+"?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		db.Close()
		return err
	}

	db = new_db
	return err
}

func setup() error {
	query := `
		CREATE TABLE IF NOT EXISTS applications (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			company TEXT NOT NULL,
			role TEXT NOT NULL,
			location TEXT NOT NULL,
			date TEXT,
			status TEXT NOT NULL
		);
	`

	_, err := db.Exec(query)
	return err
}

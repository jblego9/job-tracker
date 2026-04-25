package repository

import (
	"database/sql"
	"os"

	"backend"
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
		dbPath = "./data/dev.db"
	}

	return dbPath
}

func open() error {
	new_db, err := sql.Open("sqlite", dbPath())
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
			text TEXT
		);
	`

	_, err := db.Exec(query)
	return err
}

func GetApplications() ([]backend.Application, error) {
	query := `
		SELECT id, text FROM applications;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]backend.Application, 0)
	for rows.Next() {
		var app backend.Application
		if err := rows.Scan(&app.Id, &app.Text); err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}

func InsertApplication(app backend.Application) error {
	query := `INSERT INTO applications (text) VALUES (?)`

	_, err := db.Exec(query, app.Text)
	return err
}

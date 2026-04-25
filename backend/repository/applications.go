package repository

import (
	"backend"
	"database/sql"
)

func GetApplications() ([]backend.Application, error) {
	rows, err := db.Query("SELECT id, text FROM applications;")
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

func GetApplication(id int) (*backend.Application, error) {
	row := db.QueryRow("SELECT id, text FROM applications WHERE id = ?;", id)

	var app backend.Application
	if err := row.Scan(&app.Id, &app.Text); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return &app, nil
}

func InsertApplication(app backend.Application) error {
	_, err := db.Exec("INSERT INTO applications (text) VALUES (?);", app.Text)
	return err
}

func UpdateApplication(app backend.Application) error {
	_, err := db.Exec("UPDATE applications SET text = ? WHERE id = ?;", app.Text, app.Id)
	return err
}

func DeleteApplication(id int) error {
	_, err := db.Exec("DELETE FROM applications WHERE id = ?;", id)
	return err
}

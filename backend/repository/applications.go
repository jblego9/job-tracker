package repository

import (
	"backend"
	"database/sql"
)

func GetApplications() ([]backend.Application, error) {
	query := `SELECT id, text FROM applications;`

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

func GetApplication(id int) (*backend.Application, error) {
	query := `SELECT id, text FROM applications WHERE id = (?);`

	row := db.QueryRow(query, id)

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
	query := `INSERT INTO applications (text) VALUES (?);`

	_, err := db.Exec(query, app.Text)
	return err
}

func DeleteApplication(id int) error {
	query := `DELETE FROM applications WHERE id = (?);`

	_, err := db.Exec(query, id)
	return err
}

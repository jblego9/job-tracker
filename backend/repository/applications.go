package repository

import (
	"backend"
	"database/sql"
)

func getApplications(query string, args ...any) ([]backend.Application, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]backend.Application, 0)
	for rows.Next() {
		var app backend.Application

		if err := rows.Scan(&app.Id, &app.Company, &app.Role, &app.Location, &app.Date, &app.Status); err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}

func GetApplications() ([]backend.Application, error) {
	return getApplications("SELECT id, company, role, location, date, status FROM applications;")
}

func GetApplicationsByStatus(status string) ([]backend.Application, error) {
	return getApplications("SELECT id, company, role, location, date, status FROM applications WHERE status = ?;", status)
}

func GetApplication(id int) (*backend.Application, error) {
	row := db.QueryRow("SELECT id, company, role, location, date, status FROM applications WHERE id = ?;", id)

	var app backend.Application
	if err := row.Scan(&app.Id, &app.Company, &app.Role, &app.Location, &app.Date, &app.Status); err != nil {
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
	query := `
		INSERT INTO applications (company, role, location, date, status)
		VALUES (?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query, app.Company, app.Role, app.Location, app.Date, app.Status)
	return err
}

func UpdateApplication(app backend.Application) error {
	query := `
		UPDATE applications
		SET company = ?, role = ?, location = ?, date = ?, status = ?
		WHERE id = ?;
	`

	_, err := db.Exec(query, app.Company, app.Role, app.Location, app.Date, app.Status, app.Id)
	return err
}

func DeleteApplication(id int) error {
	_, err := db.Exec("DELETE FROM applications WHERE id = ?;", id)
	return err
}

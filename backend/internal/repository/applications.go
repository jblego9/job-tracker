package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"backend/internal"
)

func getApplications(queryAdditions []string, args ...any) ([]internal.Application, error) {
	var query strings.Builder
	query.WriteString(`SELECT id, company, role, location, date, status FROM applications`)
	for _, addition := range queryAdditions {
		query.WriteString(" " + addition)
	}
	query.WriteString(";")

	rows, err := db.Query(query.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]internal.Application, 0)
	for rows.Next() {
		var app internal.Application

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

func GetApplications(statusFilter string, sortBy string, sortDir string, pageNum *int, pageSize *int) ([]internal.Application, error) {
	queryAdditions := make([]string, 0)
	queryArguments := make([]any, 0)

	if statusFilter != "" {
		queryAdditions = append(queryAdditions, "WHERE status = ?")
		queryArguments = append(queryArguments, statusFilter)
	}

	if sortBy != "" && sortDir != "" {
		queryAdditions = append(queryAdditions, "ORDER BY", sortBy, strings.ToUpper(sortDir))
	}

	if pageNum != nil && pageSize != nil {
		queryAdditions = append(queryAdditions, "LIMIT", strconv.Itoa(*pageSize), "OFFSET", strconv.Itoa(*pageNum**pageSize))
	}

	return getApplications(queryAdditions, queryArguments...)
}

func GetApplication(id int) (*internal.Application, error) {
	row := db.QueryRow("SELECT id, company, role, location, date, status FROM applications WHERE id = ?;", id)

	var app internal.Application
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

func InsertApplication(app internal.Application) error {
	query := `
		INSERT INTO applications (company, role, location, date, status)
		VALUES (?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query, app.Company, app.Role, app.Location, app.Date, app.Status)
	return err
}

func UpdateApplication(app internal.Application) error {
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

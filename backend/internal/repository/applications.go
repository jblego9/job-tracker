package repository

import (
	"database/sql"
	"errors"
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
	exists, err := DoesApplicationExist(app.Id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Application with id " + strconv.Itoa(app.Id) + " does not exist")
	}

	query := `
		UPDATE applications
		SET company = ?, role = ?, location = ?, date = ?, status = ?
		WHERE id = ?;
	`

	_, err = db.Exec(query, app.Company, app.Role, app.Location, app.Date, app.Status, app.Id)
	return err
}

func DeleteApplication(id int) error {
	exists, err := DoesApplicationExist(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Application with id " + strconv.Itoa(id) + " does not exist")
	}

	_, err = db.Exec("DELETE FROM applications WHERE id = ?;", id)
	return err
}

func DoesApplicationExist(id int) (bool, error) {
	var ignoredId int
	err := db.QueryRow("SELECT id FROM applications WHERE id = ?;", id).Scan(&ignoredId)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func GetSummary() (*internal.Summary, error) {
	query := `SELECT COUNT(*) AS count, status FROM applications GROUP BY status`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summary internal.Summary
	statusMap := make(map[string]int)

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&count, &status); err != nil {
			return nil, err
		}

		statusMap[status] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	summary = internal.Summary{
		TotalNoResponse:        statusMap["NO_RESPONSE"],
		TotalRejected:          statusMap["REJECTED"],
		TotalAwaitingInterview: statusMap["AWAITING_INTERVIEW"],
		TotalAccepted:          statusMap["ACCEPTED"],
	}
	summary.Total = summary.TotalNoResponse + summary.TotalRejected + summary.TotalAwaitingInterview + summary.TotalAccepted

	return &summary, nil
}

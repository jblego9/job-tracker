package internal

import "strconv"

type Application struct {
	Id       int
	Company  string
	Role     string
	Location string
	Date     string
	Status   string
}

func ApplicationToCsvRecord(app Application) []string {
	return []string{
		strconv.Itoa(app.Id),
		app.Company,
		app.Role,
		app.Location,
		app.Date,
		app.Status,
	}
}

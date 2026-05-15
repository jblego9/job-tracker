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

func ValidateStatus(status string) bool {
	switch status {
	case "AWAITING_RESPONSE":
		fallthrough
	case "NO_RESPONSE":
		fallthrough
	case "REJECTED":
		fallthrough
	case "AWAITING_INTERVIEW":
		fallthrough
	case "ACCEPTED":
		return true	
	default:
		return false	
	}
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

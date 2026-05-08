package service

import "strconv"

func ValidateStatus(status string) string {
	var validated string
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
		validated = status 
	default:
		validated = ""
	}

	return validated
}

func validateSortBy(sortBy string) string {
	var validated string
	switch sortBy {
	case "id":
		fallthrough
	case "company":
		fallthrough
	case "role":
		fallthrough
	case "location":
		fallthrough
	case "date":
		validated = sortBy
	default:
		validated = ""
	}

	return validated
}

func validateSortDir(sortDir string) string {
	var validated string
	switch sortDir {
	case "asc":
		fallthrough
	case "desc":
		validated = sortDir
	default:
		validated = ""
	}

	return validated
}

func validatePageParam(param string) *int {
	var value int

	parsed, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		return nil
	}
	value = int(parsed)

	return &value
}

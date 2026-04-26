package service

func validateStatusFilter(filter string) string {
	var validated string
	switch filter {
	case "NO_RESPONSE":
		fallthrough
	case "REJECTED":
		fallthrough
	case "AWAITING_INTERVIEW":
		fallthrough
	case "ACCEPTED":
		validated = filter
	default:
		validated = ""
	}

	return validated
}

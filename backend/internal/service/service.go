package service

import (
	"backend/internal"
	"backend/internal/repository"
	"strconv"
	"fmt"
)

func GetApplications(statusFilter string, sortBy string, sortDir string, pageNum string, pageSize string) ([]internal.Application, error) {
	if !internal.ValidateStatus(statusFilter) && statusFilter != "" {
		return nil, fmt.Errorf("Invalid status '%s'", statusFilter)
	}

	if !validateSortBy(sortBy) && sortBy != "" {
		return nil, fmt.Errorf("Invalid sort by '%s'", sortBy) 
	}

	if !validateSortDir(sortDir) && sortDir != "" {
		return nil, fmt.Errorf("Invalid sort dir '%s'", sortDir)
	}


	parsedPageNum, err := parsePageParam(pageNum)
	if err != nil && pageNum != "" {
		return nil, fmt.Errorf("Invalid page num '%s'", pageNum)
	}

	parsedPageSize, err := parsePageParam(pageSize)
	if err != nil && pageSize != "" {
		return nil, fmt.Errorf("Invalid page size '%s'", pageSize)
	}

	return repository.GetApplications(
		statusFilter,
		sortBy,
		sortDir,
		parsedPageNum,
		parsedPageSize,
	)
}

func parsePageParam(param string) (*int, error) {
	var value int

	parsed, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		return nil, err
	}
	value = int(parsed)

	return &value, nil
}

func validateSortBy(sortBy string) bool {
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
		return true
	default:
		return false
	}
}

func validateSortDir(sortDir string) bool {
	switch sortDir {
	case "asc":
		fallthrough
	case "desc":
		return true
	default:
		return false
	}
}

func GetApplication(id int) (*internal.Application, error) {
	return repository.GetApplication(id)
}

func AddApplication(app internal.Application) error {
	return repository.InsertApplication(app)
}

func UpdateApplication(app internal.Application) error {
	return repository.UpdateApplication(app)
}

func DeleteApplication(id int) error {
	return repository.DeleteApplication(id)
}

func DoesApplicationExist(id int) (bool, error) {
	return repository.DoesApplicationExist(id)
}

func GetSummary() (*internal.Summary, error) {
	return repository.GetSummary()
}

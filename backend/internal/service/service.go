package service

import (
	"backend/internal"
	"backend/internal/repository"
)

func GetApplications(statusFilter string, sortBy string, sortDir string, pageNum string, pageSize string) ([]internal.Application, error) {
	return repository.GetApplications(
		validateStatusFilter(statusFilter),
		validateSortBy(sortBy),
		validateSortDir(sortDir),
		validatePageParam(pageNum),
		validatePageParam(pageSize),
	)
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

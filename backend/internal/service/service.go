package service

import (
	"backend/internal"
	"backend/internal/repository"
)

// TODO: See if SQL has a feature to check for a row without getting it
func DoesApplicationExist(id int) (bool, error) {
	app, err := GetApplication(id)
	if err != nil || app == nil {
		return false, err
	}

	return true, err
}

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

// TODO: Find a better way to get the status counts. Maybe SQL has a feature for it?
func GetSummary() (*internal.Summary, error) {
	var summary internal.Summary

	apps, err := repository.GetApplications("NO_RESPONSE", "", "", nil, nil)
	if err != nil {
		return nil, err
	}
	summary.TotalNoResponse = len(apps)

	apps, err = repository.GetApplications("REJECTED", "", "", nil, nil)
	if err != nil {
		return nil, err
	}
	summary.TotalRejected = len(apps)

	apps, err = repository.GetApplications("AWAITING_INTERVIEW", "", "", nil, nil)
	if err != nil {
		return nil, err
	}
	summary.TotalAwaitingInterview = len(apps)

	apps, err = repository.GetApplications("ACCEPTED", "", "", nil, nil)
	if err != nil {
		return nil, err
	}
	summary.TotalAccepted = len(apps)

	summary.Total = summary.TotalNoResponse + summary.TotalRejected + summary.TotalAwaitingInterview + summary.TotalAccepted

	return &summary, nil
}

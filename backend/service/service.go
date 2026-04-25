package service

import (
	"backend"
	"backend/repository"
)

func DoesApplicationExist(id int) (bool, error) {
	app, err := GetApplication(id)
	if err != nil || app == nil {
		return false, err
	}

	return true, err
}

func GetApplications() ([]backend.Application, error) {
	return repository.GetApplications()
}

func GetApplication(id int) (*backend.Application, error) {
	return repository.GetApplication(id)
}

func AddApplication(app backend.Application) error {
	return repository.InsertApplication(app)
}

func UpdateApplication(app backend.Application) error {
	return repository.UpdateApplication(app)
}

func DeleteApplication(id int) error {
	return repository.DeleteApplication(id)
}

func GetSummary() (*backend.Summary, error) {
	var summary backend.Summary

	apps, err := repository.GetApplicationsByStatus("NO_RESPONSE")
	if err != nil {
		return nil, err
	}
	summary.TotalNoResponse = len(apps)

	apps, err = repository.GetApplicationsByStatus("REJECTED")
	if err != nil {
		return nil, err
	}
	summary.TotalRejected = len(apps)

	apps, err = repository.GetApplicationsByStatus("AWAITING_INTERVIEW")
	if err != nil {
		return nil, err
	}
	summary.TotalAwaitingInterview = len(apps)

	apps, err = repository.GetApplicationsByStatus("ACCEPTED")
	if err != nil {
		return nil, err
	}
	summary.TotalAccepted = len(apps)

	summary.Total = summary.TotalNoResponse + summary.TotalRejected + summary.TotalAwaitingInterview + summary.TotalAccepted

	return &summary, nil
}

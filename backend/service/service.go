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

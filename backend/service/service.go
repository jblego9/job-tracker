package service

import (
	"backend"
	"backend/repository"
)

func GetApplications() ([]backend.Application, error) {
	return repository.GetApplications()
}

func AddApplication(app backend.Application) error {
	return repository.InsertApplication(app)
}

func DeleteApplication(id uint) error {
	return repository.DeleteApplication(id)
}

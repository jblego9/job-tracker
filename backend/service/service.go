package service

import (
	backend "backend"
	"backend/repository"
)

func GetApplications() ([]backend.Application, error) {
	return repository.GetApplications()
}

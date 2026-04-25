package controller

import (
	"backend"
	"encoding/json"
	"net/http"

	"backend/service"
)

func GetApplications(w http.ResponseWriter, _ *http.Request) {
	apps, err := service.GetApplications()
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	writeJsonResponse(w, apps)
}

func AddApplication(w http.ResponseWriter, r *http.Request) {
	var app backend.Application
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	if err := service.AddApplication(app); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

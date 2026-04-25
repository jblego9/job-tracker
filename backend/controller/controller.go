package controller

import (
	"backend"
	"encoding/json"
	"net/http"
	"strconv"

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

	w.WriteHeader(http.StatusCreated)
}

func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	parsed_id, err := strconv.ParseInt(r.PathValue("id"), 0, 64)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	id := int(parsed_id)
	if err := service.DeleteApplication(id); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

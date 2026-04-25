package controller

import (
	"backend"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/service"
)

func parsePathId(r *http.Request) (int, error) {
	parsed_id, err := strconv.ParseInt(r.PathValue("id"), 0, 64)
	if err != nil {
		return 0, err
	}

	return int(parsed_id), err
}

func GetApplications(w http.ResponseWriter, _ *http.Request) {
	apps, err := service.GetApplications()
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	writeJsonResponse(w, apps)
}

func GetApplication(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathId(r)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	app, err := service.GetApplication(id)
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	if app == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeJsonResponse(w, app)
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
	id, err := parsePathId(r)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	app, err := service.GetApplication(id)
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	if app == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := service.DeleteApplication(id); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

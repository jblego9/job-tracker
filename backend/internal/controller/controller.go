package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal"
	"backend/internal/service"
)

func parsePathId(r *http.Request) (int, error) {
	parsedId, err := strconv.ParseInt(r.PathValue("id"), 0, 64)
	if err != nil {
		return 0, err
	}

	return int(parsedId), err
}

func GetApplications(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	statusFilter := params.Get("statusFilter")

	apps, err := service.GetApplications(statusFilter)
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
	var app internal.Application
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

func UpdateApplication(w http.ResponseWriter, r *http.Request) {
	var updated_app internal.Application
	if err := json.NewDecoder(r.Body).Decode(&updated_app); err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	exists, err := service.DoesApplicationExist(updated_app.Id)
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := service.UpdateApplication(updated_app); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathId(r)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	exists, err := service.DoesApplicationExist(id)
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := service.DeleteApplication(id); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetSummary(w http.ResponseWriter, _ *http.Request) {
	summary, err := service.GetSummary()
	if err != nil || summary == nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	writeJsonResponse(w, *summary)
}

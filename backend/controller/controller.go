package controller

import (
	"backend"
	"encoding/json"
	"io"
	"net/http"

	"backend/service"
)

func writeInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}

func writeBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, err.Error())
}

func writeJsonResponse(w http.ResponseWriter, v any) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

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

package controller

import (
	"backend/service"
	"encoding/json"
	"io"
	"net/http"
)

func writeInternalErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}

func writeJsonResponse(w http.ResponseWriter, v any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func GetApplications(w http.ResponseWriter, _ *http.Request) {
	apps, err := service.GetApplications()
	if err != nil {
		writeInternalErrorResponse(w, err)
		return
	}

	writeJsonResponse(w, apps)
}

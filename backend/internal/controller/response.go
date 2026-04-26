package controller

import (
	"encoding/json"
	"io"
	"net/http"
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

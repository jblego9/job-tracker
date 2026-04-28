package controller

import (
	"encoding/csv"
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

// TODO: Prepare content before writing
func writeJsonResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}
}

// TODO: Prepare content before writing
func writeCsvResponse(w http.ResponseWriter, records [][]string) {
	w.Header().Set("Content-Type", "text/csv")
	if err := csv.NewWriter(w).WriteAll(records); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}
}

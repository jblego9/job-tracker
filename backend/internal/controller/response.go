package controller

import (
	"bytes"
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

func writeJsonResponse(w http.ResponseWriter, v any) {
	bytes, err := json.Marshal(v)
	if err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func writeCsvResponse(w http.ResponseWriter, records [][]string) {
	var buffer bytes.Buffer
	if err := csv.NewWriter(&buffer).WriteAll(records); err != nil {
		writeInternalServerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Write(buffer.Bytes())
}

package server

import (
	"fmt"
	"net/http"
	"os"

	"backend/internal/controller"
	"backend/internal/middleware"
)

func Start() {
	setup()

	port := port()
	fmt.Printf("Server listening on \"localhost:%s\"...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

func setup() {
	http.HandleFunc("GET /api/v1/applications", middleware.Logging(controller.GetApplications))
	http.HandleFunc("GET /api/v1/applications/{id}", middleware.Logging(controller.GetApplication))
	http.HandleFunc("POST /api/v1/applications", middleware.Logging(controller.AddApplication))
	http.HandleFunc("PUT /api/v1/applications", middleware.Logging(controller.UpdateApplication))
	http.HandleFunc("DELETE /api/v1/applications/{id}", middleware.Logging(controller.DeleteApplication))
	http.HandleFunc("GET /api/v1/applications/summary", middleware.Logging(controller.GetSummary))
	http.HandleFunc("GET /api/v1/applications/csv-export", middleware.Logging(controller.GetApplicationsCsvExport))
}

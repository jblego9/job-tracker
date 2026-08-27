package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"backend/internal"
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/service"
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

func noResponseCheckAfterMonths() error {
	envValue := os.Getenv("NO_RESPONSE_CHECK_AFTER_N_MONTHS")
	if envValue == "" {
		return nil
	}

	months, err := strconv.Atoi(envValue)
	if err != nil {
		return err
	}

	return service.NoResponseCheck(months)
}

func setup() {
	subFS, err := fs.Sub(internal.EmbeddedFiles, "frontend-build")
	if err != nil {
		panic(err)
	}

	err = noResponseCheckAfterMonths()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("GET /api/v1/applications", middleware.Logging(controller.GetApplications))
	http.HandleFunc("GET /api/v1/applications/{id}", middleware.Logging(controller.GetApplication))
	http.HandleFunc("POST /api/v1/applications", middleware.Logging(controller.AddApplication))
	http.HandleFunc("PUT /api/v1/applications", middleware.Logging(controller.UpdateApplication))
	http.HandleFunc("DELETE /api/v1/applications/{id}", middleware.Logging(controller.DeleteApplication))
	http.HandleFunc("GET /api/v1/applications/summary", middleware.Logging(controller.GetSummary))
	http.HandleFunc("GET /api/v1/applications/csv-export", middleware.Logging(controller.GetApplicationsCsvExport))

	http.Handle("/", http.FileServer(http.FS(subFS)))
}

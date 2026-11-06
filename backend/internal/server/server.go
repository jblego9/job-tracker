package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"backend/internal"
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/service"
)

func Start() {
	setup()

	port := port()
	fmt.Printf("Server listening on \"localhost:%s\"...\n", port)

	err := noResponseCheckAfterMonths()
	if err != nil {
		panic(err)
	}

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

	fmt.Printf("Performing no response after '%d' months check...\n", months)
	return service.NoResponseCheck(months)
}

func singlePageAppHandler(fsys fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(fsys))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Unknown api route
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		// Relative filesystem path
		name := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if name == "" {
			name = "."
		}

		// Serve files
		if info, err := fs.Stat(fsys, name); err == nil && !info.IsDir() {
			fileServer.ServeHTTP(w, r)
			return
		}

		// React router handles everything else
		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	})
}

func setup() {
	subFS, err := fs.Sub(internal.EmbeddedFiles, "frontend-build")
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
	http.HandleFunc("GET /api/v1/applications/history", middleware.Logging(controller.GetHistory))
	http.HandleFunc("GET /api/v1/applications/history/{year}", middleware.Logging(controller.GetHistory))

	http.Handle("/", singlePageAppHandler(subFS))
}

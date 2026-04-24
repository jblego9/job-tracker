package server

import (
	"fmt"
	"net/http"
	"os"

	"backend/controller"
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
	http.HandleFunc("/api/v1/applications", controller.GetApplications)
}

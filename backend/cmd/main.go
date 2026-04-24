package main

import (
	"backend/controller"
	"backend/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	godotenv.Load()

	if err := repository.Init(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/v1/applications", controller.GetApplications)

	portString := os.Getenv("PORT")

	fmt.Printf("Server listening on 'localhost:%s'...\n", portString)
	http.ListenAndServe(fmt.Sprintf(":%s", portString), nil)
}

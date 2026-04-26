package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	"backend/internal/repository"
)

func main() {
	godotenv.Load()

	if err := repository.Init(); err != nil {
		log.Fatal(err)
	}

	apps := GetTestData()
	for _, app := range apps {
		if err := repository.InsertApplication(app); err != nil {
			log.Fatal(err)
		}
	}
}

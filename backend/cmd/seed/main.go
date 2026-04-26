package main

import (
	"backend/repository"
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
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

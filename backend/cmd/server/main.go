package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	"backend/internal/repository"
	"backend/internal/server"
)

func main() {
	godotenv.Load()

	if err := repository.Init(); err != nil {
		log.Fatal(err)
	}

	server.Start()
}

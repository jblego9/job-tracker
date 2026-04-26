package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	"backend/repository"
	"backend/server"
)

func main() {
	godotenv.Load()

	if err := repository.Init(); err != nil {
		log.Fatal(err)
	}

	server.Start()
}

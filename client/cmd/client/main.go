package main

import (
	"log"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/config"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/routes"
)

func main() {
	config.LoadEnv()

	config.LoadDatabaseConnection()

	router := routes.SetupRouter()

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

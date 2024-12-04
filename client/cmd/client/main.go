package main

import (
	"log"
	"os"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/config"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/routes"
)

func main() {
	config.LoadEnv()

	config.LoadDatabaseConnection()

	router := routes.SetupRouter()

	if err := router.Run(os.Getenv("GIN_PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

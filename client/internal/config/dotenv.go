package config

import (
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/pkg/utils"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		utils.Logger.Println("Encountered an error during .env load")
	}
}

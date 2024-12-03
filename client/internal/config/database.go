package config

import (
	"fmt"
	"os"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabaseConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=%s TimeZone=%s",
		username, password, database, port, host, "disable", "America/Sao_Paulo")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatal("Encountered an error during trying stablish connection with database: ", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		utils.Logger.Fatal("Failed to retrieve underlying SQL DB: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		utils.Logger.Fatal("Failed to ping database: ", err)
	}

	utils.Logger.Print("Database connection established successfully")

	return db
}

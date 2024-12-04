package config

import (
	"fmt"
	"os"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/models"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	databaseInstance *gorm.DB
)

func LoadDatabaseConnection() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=%s TimeZone=%s",
		username, password, database, port, host, "disable", "America/Sao_Paulo")

	var err error

	databaseInstance, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		utils.Logger.Fatal("Encountered an error during trying to establish connection with database: ", err)
	}

	sqlDB, err := databaseInstance.DB()

	if err != nil {
		utils.Logger.Fatal("Failed to retrieve underlying SQL DB: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		utils.Logger.Fatal("Failed to ping database: ", err)
	}

	if err := databaseInstance.AutoMigrate(&models.Transaction{}); err != nil {
		utils.Logger.Fatal("Error during database migration: ", err)
	}

	utils.Logger.Print("Database connection established successfully")
}

func GetDatabaseConnection() *gorm.DB {
	if databaseInstance == nil {
		utils.Logger.Fatal("Database connection is not initialised. Call `LoadDatabaseConnection` first.")
	}

	return databaseInstance
}

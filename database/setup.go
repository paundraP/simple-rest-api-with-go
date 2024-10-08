package database

import (
	"log"
	"os"
	"paundraP/rest-api-with-go/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func ConnectDB() {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		log.Fatal("Database environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	err = db.AutoMigrate(&models.Ticket{})
	if err != nil {
		log.Fatal("Failed to auto-migrate Ticket model")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate Ticket model")
	}
	DB = db
}

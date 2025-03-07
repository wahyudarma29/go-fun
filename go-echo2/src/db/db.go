package db

import (
	"fmt"
	"go-echo2/src/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
	log.Println("Database connected successfully!")
}

func TestConnection() {
	Connect()
	// if we want to close connection we need to access the base sql instances
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB instance:", err)
	} else {
		log.Println("DB Connection Success")
	}
	defer sqlDB.Close()
}
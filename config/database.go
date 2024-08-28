package config

import (
	"fmt"
	"os"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	"slc2/models"
)

func InitDB()*gorm.DB{
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword,dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Melakukan auto-migration untuk membuat tabel berdasarkan model yang ada
	db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Rental{})
	// DB = db
	log.Println("Database Connection Established Successfully")
	return db
}

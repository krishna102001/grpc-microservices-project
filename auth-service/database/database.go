package database

import (
	"fmt"
	"os"

	logger "github.com/charmbracelet/log"
	"github.com/krishna102001/grpc-microservices-project/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if host == "" || username == "" || password == "" || name == "" || port == "" {
		logger.Fatal("Database Env file not found")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect the database %v", err)
	}
	logger.Info("Database is connected Successfully")
	DB = db

	migrate()
}

func migrate() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		logger.Fatal("Failed to migrate the model")
	}
}

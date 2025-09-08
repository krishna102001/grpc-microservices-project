package database

import (
	"fmt"
	"log"
	"os"

	"github.com/krishna102001/grpc-microservices-project/blog-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if host == "" || username == "" || password == "" || port == "" || name == "" {
		log.Fatal("Failed to get the env value")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with blog database")
	}

	DB = db
	log.Println("Successfully Blog Database Connected")

	migrate()
}

func migrate() {
	if err := DB.AutoMigrate(&model.Blog{}); err != nil {
		log.Fatalf("Failed to migrate the blog database %v", err)
	}
	log.Println("Migration of Blog Services Done Successfull")
}

package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/krishna102001/grpc-microservices-project/blog-service/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env file")
	}
	database.Connect()
}

func main() {

}

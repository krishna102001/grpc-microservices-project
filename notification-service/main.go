package main

import (
	"log"
	"os"
	"sync"

	"github.com/krishna102001/grpc-microservices-project/notification-service/consumer"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalf("Failed to load the env file %v", err)
// 	}
// }

func main() {
	brk_url := os.Getenv("BROKER_URL")
	tpc := os.Getenv("TOPIC")
	if brk_url == "" || tpc == "" {
		log.Fatal("either broker or topic not found in env file")
	}

	brokerURL := []string{brk_url}

	var wg sync.WaitGroup
	if err := consumer.Connect(brokerURL, tpc, &wg); err != nil {
		log.Fatal("Failed to connect with kafka")
	}

	log.Println("Successfully connected to kafka")
	wg.Wait()
}

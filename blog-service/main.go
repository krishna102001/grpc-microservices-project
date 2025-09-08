package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/krishna102001/grpc-microservices-project/blog-service/database"
	"github.com/krishna102001/grpc-microservices-project/blog-service/pb"
	"github.com/krishna102001/grpc-microservices-project/blog-service/service"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env file")
	}
	database.Connect()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env value not found")
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to connect to tcp server")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBlog_ServiceServer(grpcServer, &service.BlogServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the gRPC Server")
	}
	log.Printf("gRPC server started of Blog Service on port on %v", lis.Addr())
}

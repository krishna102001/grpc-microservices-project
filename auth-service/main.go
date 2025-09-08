package main

import (
	"net"
	"os"

	logger "github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/krishna102001/grpc-microservices-project/auth-service/pb"
	"github.com/krishna102001/grpc-microservices-project/auth-service/services"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Error("Failed to load the env file")
	}
}

func main() {
	//get port no.
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatalf("Port is empty")
	}

	// start TCP server
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatalf("Failed to listen on port %v", port)
	}

	// create gRPC server
	grpcServer := grpc.NewServer()

	// Register services
	authServer := &services.AuthServer{}
	pb.RegisterAuth_ServiceServer(grpcServer, authServer)

	// start gRPC server
	logger.Infof("gRPC server is started successfully on port no. %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve the grpc server on port no. %v", lis.Addr())
	}

}

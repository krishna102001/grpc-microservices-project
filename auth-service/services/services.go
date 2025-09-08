package services

import "github.com/krishna102001/grpc-microservices-project/auth-service/pb"

// AuthServer struct implementing the gRPC Interfaces
type AuthServer struct {
	pb.UnimplementedAuth_ServiceServer
}

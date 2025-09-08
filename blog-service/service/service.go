package service

import "github.com/krishna102001/grpc-microservices-project/blog-service/pb"

type BlogServer struct {
	pb.UnimplementedBlog_ServiceServer
}

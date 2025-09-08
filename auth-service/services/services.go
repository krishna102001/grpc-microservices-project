package services

import (
	"context"
	"fmt"

	logger "github.com/charmbracelet/log"
	"github.com/krishna102001/grpc-microservices-project/auth-service/database"
	"github.com/krishna102001/grpc-microservices-project/auth-service/models"
	"github.com/krishna102001/grpc-microservices-project/auth-service/pb"
	"github.com/krishna102001/grpc-microservices-project/auth-service/utils"
)

// AuthServer struct implementing the gRPC Interfaces
type AuthServer struct {
	pb.UnimplementedAuth_ServiceServer
}

func (srv *AuthServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.AuthResponse, error) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Warn("Failed to hash the password")
		return nil, err
	}

	user := models.User{Name: req.Name, Email: req.Email, Password: hashPassword}

	if err := database.DB.Create(&user).Error; err != nil {
		logger.Errorf("Failed to insert the data in database %v", err)
		return nil, err
	}

	token, err := utils.GenerateToken(req.Email)
	if err != nil {
		logger.Warn("Failed to Generate the Token")
		return nil, err
	}

	return &pb.AuthResponse{Token: token, Message: "User Signup Successfully"}, nil
}

func (srv *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	var user models.User

	if err := database.DB.Where("email=?", req.Email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if !utils.ComparePassword(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	token, err := utils.GenerateToken(req.Email)
	if err != nil {
		logger.Warn("Failed to Generate the Token")
		return nil, err
	}
	return &pb.AuthResponse{Token: token, Message: "User Login Successfully"}, nil
}

func (srv *AuthServer) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordRequest) (*pb.ForgetPasswordResponse, error) {
	return &pb.ForgetPasswordResponse{Message: "Successfully Reset the password"}, nil
}

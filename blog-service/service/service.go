package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/krishna102001/grpc-microservices-project/blog-service/database"
	"github.com/krishna102001/grpc-microservices-project/blog-service/model"
	"github.com/krishna102001/grpc-microservices-project/blog-service/pb"
	"github.com/krishna102001/grpc-microservices-project/blog-service/producer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BlogServer struct {
	pb.UnimplementedBlog_ServiceServer
}

type kafkaMessage struct {
	ServiceType    string `json:"service_type"`
	MessageType    string `json:"message_type"`
	MessageContent string `json:"MessageContent"`
}

func (srv *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.BlogResponse, error) {
	blog := model.Blog{Title: req.Title, Description: req.Description, AuthorId: req.AuthorId}

	if err := database.DB.Create(&blog).Error; err != nil {
		return nil, fmt.Errorf("blog Service Internal server error %v", err.Error())
	}

	tpc := os.Getenv("TOPIC")

	if tpc == "" {
		log.Printf("------------------ failed to get the topic name-----------------")
	}

	kf := &kafkaMessage{
		ServiceType:    "email",
		MessageType:    "create_blog",
		MessageContent: "You have Successfully created blog",
	}

	msg, err := json.Marshal(kf)
	if err != nil {
		log.Printf("---------------------- failed to marshal the json --------------------")
	}

	if err := producer.PushToKafkaQueue(tpc, msg); err != nil {
		log.Printf("------------------------- failed to send the notification ----------------- %v", err)
	}

	return &pb.BlogResponse{
		Id:          blog.Id.String(),
		Title:       blog.Title,
		Description: blog.Description,
		AuthorId:    blog.AuthorId,
	}, nil
}

func (srv *BlogServer) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.BlogListResponse, error) {
	var blogs []model.Blog
	if err := database.DB.Where("author_id = ?", req.Id).Find(&blogs).Error; err != nil {
		return nil, fmt.Errorf("blog service internal server error %v", err.Error())
	}

	var blogResponses []*pb.BlogResponse
	for _, blog := range blogs {
		blogResponses = append(blogResponses, &pb.BlogResponse{
			Id:          blog.Id.String(),
			Title:       blog.Title,
			Description: blog.Description,
			AuthorId:    blog.AuthorId,
		})
	}
	return &pb.BlogListResponse{
		Blogs: blogResponses,
	}, nil
}

func (srv *BlogServer) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.BlogResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	emailList := md.Get("email")
	if len(emailList) == 0 {
		log.Println("unauthorized: missing user ID in metadata")
		return nil, status.Error(codes.Unauthenticated, "unauthorized: missing user ID in metadata")
	}

	userID := emailList[0]

	var existingBlog model.Blog
	if err := database.DB.Where("id = ? AND author_id = ?", req.Id, userID).First(&existingBlog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("blog not found or you are not authorized to update it")
		}
		return nil, fmt.Errorf("database error while fetching blog: %v", err)
	}

	updateData := make(map[string]interface{})
	if req.Title != "" {
		updateData["Title"] = req.Title
	}
	if req.Description != "" {
		updateData["Description"] = req.Description
	}
	if err := database.DB.Model(&model.Blog{}).Where("id = ? AND author_id = ?", req.Id, userID).Updates(updateData).Error; err != nil {
		return nil, fmt.Errorf("blog service internal server error %v", err.Error())
	}

	if err := database.DB.Where("id = ?", req.Id).First(&existingBlog).Error; err != nil {
		return nil, fmt.Errorf("blog service internal server error")
	}

	tpc := os.Getenv("TOPIC")

	if tpc == "" {
		log.Printf("------------------ failed to get the topic name-----------------")
	}

	kf := &kafkaMessage{
		ServiceType:    "email",
		MessageType:    "update_blog",
		MessageContent: "You have Successfully updated blog",
	}

	msg, err := json.Marshal(kf)
	if err != nil {
		log.Printf("---------------------- failed to marshal the json --------------------")
	}

	if err := producer.PushToKafkaQueue(tpc, msg); err != nil {
		log.Printf("------------------------- failed to send the notification ----------------- %v", err)
	}

	return &pb.BlogResponse{
		Id:          existingBlog.Id.String(),
		Title:       existingBlog.Title,
		Description: existingBlog.Description,
		AuthorId:    existingBlog.AuthorId,
	}, nil
}

func (srv *BlogServer) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.MessageResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	emailList := md.Get("email")
	if len(emailList) == 0 {
		return nil, status.Error(codes.NotFound, "user id not found")
	}
	userId := emailList[0]
	var existingBlog model.Blog
	if err := database.DB.Where("id =? AND author_id = ?", req.Id, userId).First(&existingBlog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("blog not found or you are not authorised to delete the data")
		}
		return nil, fmt.Errorf("blog service have internal server error")
	}

	if err := database.DB.Where("id = ?", req.Id).Delete(&model.Blog{}).Error; err != nil {
		return nil, fmt.Errorf("blog not found %v", err.Error())
	}

	tpc := os.Getenv("TOPIC")

	if tpc == "" {
		log.Printf("------------------ failed to get the topic name-----------------")
	}

	kf := &kafkaMessage{
		ServiceType:    "email",
		MessageType:    "delete_blog",
		MessageContent: "You have Successfully deleted blog",
	}

	msg, err := json.Marshal(kf)
	if err != nil {
		log.Printf("---------------------- failed to marshal the json --------------------")
	}

	if err := producer.PushToKafkaQueue(tpc, msg); err != nil {
		log.Printf("------------------------- failed to send the notification ----------------- %v", err)
	}
	return &pb.MessageResponse{
		Message: "Succesfully deleted the blog",
	}, nil
}

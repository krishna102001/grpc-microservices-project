package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/krishna102001/grpc-microservices-project/blog-service/database"
	"github.com/krishna102001/grpc-microservices-project/blog-service/model"
	"github.com/krishna102001/grpc-microservices-project/blog-service/pb"
	"gorm.io/gorm"
)

type BlogServer struct {
	pb.UnimplementedBlog_ServiceServer
}

func (srv *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.BlogResponse, error) {
	blog := model.Blog{Title: req.Title, Description: req.Description, AuthorId: req.AuthorId}

	if err := database.DB.Create(&blog).Error; err != nil {
		return nil, fmt.Errorf("blog Service Internal server error %v", err.Error())
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
	var existingBlog model.Blog
	if err := database.DB.Where("id = ? AND author_id = ?", req.Id, ctx.Value("user_id")).First(&existingBlog).Error; err != nil {
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
	if err := database.DB.Where("id = ? AND author_id = ?", req.Id, ctx.Value("user_id")).Updates(updateData).Error; err != nil {
		return nil, fmt.Errorf("blog service internal server error %v", err.Error())
	}

	if err := database.DB.Where("id = ?", req.Id).First(&existingBlog).Error; err != nil {
		return nil, fmt.Errorf("blog service internal server error")
	}

	return &pb.BlogResponse{
		Id:          existingBlog.Id.String(),
		Title:       existingBlog.Title,
		Description: existingBlog.Description,
		AuthorId:    existingBlog.AuthorId,
	}, nil
}

func (srv *BlogServer) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.MessageResponse, error) {

	var existingBlog model.Blog
	if err := database.DB.Where("id =? AND author_id = ?", req.Id, ctx.Value("user_id")).First(&existingBlog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("blog not found or you are not authorised to delete the data")
		}
		return nil, fmt.Errorf("blog service have internal server error")
	}

	if err := database.DB.Where("id = ?", req.Id).Delete(&model.Blog{}).Error; err != nil {
		return nil, fmt.Errorf("blog not found %v", err.Error())
	}
	return &pb.MessageResponse{
		Message: "Succesfully deleted the blog",
	}, nil
}

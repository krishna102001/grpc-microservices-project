package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type BlogHandler struct {
	blogClient pb.Blog_ServiceClient
}

func NewBlogHandler(blogAddr string) *BlogHandler {
	conn, err := grpc.NewClient(blogAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect the Blog gRPC handler %v", err)
	}
	client := pb.NewBlog_ServiceClient(conn)
	return &BlogHandler{blogClient: client}
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var req pb.CreateBlogRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or author_id"})
		return
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	res, err := h.blogClient.CreateBlog(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"blog": res,
	})
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	var req pb.GetBlogRequest

	id, _ := c.Get("email")
	req.Id = id.(string)
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	res, err := h.blogClient.GetBlog(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var req pb.UpdateBlogRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or author_id"})
		return
	}
	id, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: email missing in context"})
		return
	}
	email, ok := id.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: invalid email format"})
		return
	}
	log.Println("email-----------------------------", email)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := h.blogClient.UpdateBlog(metadata.AppendToOutgoingContext(ctx, "email", email), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error" + err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"blog": res})
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	var req pb.DeleteBlogRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or author_id"})
		return
	}

	id, _ := c.Get("email")
	email := id.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := h.blogClient.DeleteBlog(metadata.AppendToOutgoingContext(ctx, "email", email), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": res.Message})
}

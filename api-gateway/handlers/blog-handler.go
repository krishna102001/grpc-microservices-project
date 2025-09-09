package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or author_id"})
		return
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	res, err := h.blogClient.GetBlog(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"blogs": res,
	})
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var req pb.UpdateBlogRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or author_id"})
		return
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	res, err := h.blogClient.UpdateBlog(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
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
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	res, err := h.blogClient.DeleteBlog(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "blog service internal error"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": res.Message})
}

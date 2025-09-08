package handlers

import (
	"log"

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

}

func (h *BlogHandler) GetBlog(c *gin.Context) {

}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {

}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {

}

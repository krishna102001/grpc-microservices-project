package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/handlers"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/log"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/middlewares"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load the env file")
	}
}

func main() {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler("localhost:9000")
	blogHandler := handlers.NewBlogHandler("localhost:9001")

	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)

	blogRoutes := r.Group("/blog", middlewares.AuthHandler())
	{
		blogRoutes.POST("/", blogHandler.CreateBlog)
		blogRoutes.GET("/", blogHandler.GetBlog)
		blogRoutes.PATCH("/:id", blogHandler.UpdateBlog)
		blogRoutes.DELETE("/:id", blogHandler.DeleteBlog)
	}

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed start the api gateways server")
		panic(err)
	}
	log.Info("Server started at port no. 8080")
}

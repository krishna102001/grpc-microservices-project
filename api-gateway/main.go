package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/handlers"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/log"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/middlewares"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Failed to load the env file")
// 	}
// }

func main() {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler("auth-service:9000")
	blogHandler := handlers.NewBlogHandler("blog-service:9001")

	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)
	r.POST("/forget-password", authHandler.ForgetPassword)

	blogRoutes := r.Group("/blog", middlewares.AuthHandler())
	{
		blogRoutes.POST("/", blogHandler.CreateBlog)
		blogRoutes.GET("/", blogHandler.GetBlog)
		blogRoutes.PUT("/update", blogHandler.UpdateBlog)
		blogRoutes.DELETE("/", blogHandler.DeleteBlog)
	}

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed start the api gateways server")
		panic(err)
	}
	log.Info("Server started at port no. 8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/handlers"
	"github.com/krishna102001/grpc-microservices-project/api-gateway/log"
)

func main() {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler("localhost:9000")

	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed start the api gateways server")
		panic(err)
	}
	log.Info("Server started at port no. 8080")
}

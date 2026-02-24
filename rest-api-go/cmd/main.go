package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dhion/rest-api-go/internal/handler"
	"github.com/dhion/rest-api-go/internal/repository"
	"github.com/dhion/rest-api-go/internal/service"
)

func main() {
	r := gin.Default()

	// Dependency Injection
	repo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	// Auth routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.Run(":8080")
}
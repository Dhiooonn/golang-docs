package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dhion/rest-api-go/internal/handler"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUserByID)

	r.Run(":8080")
}
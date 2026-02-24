package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/dhion/rest-api-go/internal/model"
)

var users = []model.User{
	{ID: 1, Username: "Dhion", Email: "dhion@gmail.com"},
}

// GET /users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// POST /users
func CreateUser(c *gin.Context) {
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil {

		// Return 400 kalau JSON tidak valid.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser) // Tambah user ke slice

	c.JSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {

	// GET /users/:id
	idParam := c.Param("id")

	// Convert string → int.
	id, _ := strconv.Atoi(idParam)

	// Loop cari user dengan ID yang sesuai
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
}
package model

// Entity (untuk storage / domain)
type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

// DTO Register
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// DTO Login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
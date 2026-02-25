package repository

import "github.com/dhion/rest-api-go/internal/model"

type UserRepository interface {
	GetAll() []model.User
	Create(user model.User) model.User
	GetByID(id int) (model.User, bool)
	GetByEmail(email string) (model.User, bool)
	Update(id int, user model.User) (model.User, bool)
	Delete(id int) bool
}
package repository

import "github.com/dhion/rest-api-go/internal/model"

type InMemoryUserRepository struct {
	users []model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}

func (r *InMemoryUserRepository) GetAll() []model.User {
	return r.users
}

func (r *InMemoryUserRepository) Create(user model.User) model.User {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user
}

func (r *InMemoryUserRepository) GetByID(id int) (model.User, bool) {
	for _, user := range r.users {
		if user.ID == id {
			return user, true
		}
	}
	return model.User{}, false
}

func (r *InMemoryUserRepository) GetByEmail(email string) (model.User, bool) {
	for _, user := range r.users {
		if user.Email == email {
			return user, true
		}
	}
	return model.User{}, false
}

func (r *InMemoryUserRepository) Update(id int, updatedUser model.User) (model.User, bool) {
	for i, user := range r.users {
		if user.ID == id {
			r.users[i].Username = updatedUser.Username
			r.users[i].Email = updatedUser.Email
			return r.users[i], true
		}
	}
	return model.User{}, false
}

func (r *InMemoryUserRepository) Delete(id int) bool {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true
		}
	}
	return false
}
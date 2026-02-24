package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/dhion/rest-api-go/internal/utils"

	"github.com/dhion/rest-api-go/internal/model"
	"github.com/dhion/rest-api-go/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req model.RegisterRequest) (model.User, error) {
	_, found := s.repo.GetByEmail(req.Email)
	if found {
		return model.User{}, errors.New("email sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, errors.New("gagal hash password")
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user), nil
}

func (s *UserService) Login(req model.LoginRequest) (string, error) {

	user, found := s.repo.GetByEmail(req.Email)
	if !found {
		return "", errors.New("email tidak ditemukan")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("password salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", errors.New("gagal generate token")
	}

	return token, nil
}
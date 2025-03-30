package services

import (
	"volley/internal/models"
	"volley/internal/repositories"
)

type UserService interface {
	GetAll() ([]models.User, error)
	Create(req models.CreateUserRequest) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) Create(req models.CreateUserRequest) (models.User, error) {
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
	}
	return s.repo.CreateUser(user)
}

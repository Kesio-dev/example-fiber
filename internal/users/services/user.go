package services

import (
	"airdrop/internal/users/models"
	"airdrop/internal/users/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

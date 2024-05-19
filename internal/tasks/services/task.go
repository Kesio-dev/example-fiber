package services

import (
	"airdrop/internal/repositories"
	"airdrop/internal/tasks/models"
	"airdrop/pkg/storage"
)

type TaskService struct {
	repo  *repositories.TaskRepository
	redis storage.RedisClient
}

func NewTaskService(repo *repositories.TaskRepository, redis storage.RedisClient) *TaskService {
	return &TaskService{repo: repo, redis: redis}
}

func (s *TaskService) GetByID(id int) (*models.TaskModel, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) Create(user *models.TaskModel) error {
	return s.repo.Create(user)
}

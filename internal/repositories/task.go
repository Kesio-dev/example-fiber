package repositories

import (
	"airdrop/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) FindByID(id int) (*models.TaskModel, error) {
	return nil, nil
}

func (r *TaskRepository) Create(user *models.TaskModel) error {
	return nil
}

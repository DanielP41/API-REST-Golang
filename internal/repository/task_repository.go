package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/user/argolang/internal/models"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	tasks := []models.Task{}
	err := r.db.Select(&tasks, "SELECT * FROM tasks")
	return tasks, err
}

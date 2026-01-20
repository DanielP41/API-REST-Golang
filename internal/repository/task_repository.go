package repository

import (
	"github.com/DanielP41/argolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	tasks := []models.Task{}
	err := r.db.Select(&tasks, "SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at DESC")
	return tasks, err
}

func (r *TaskRepository) Create(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetByID(id int64) (*models.Task, error) {
	task := &models.Task{}
	err := r.db.Get(task, "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1", id)
	return task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	return err
}

func (r *TaskRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

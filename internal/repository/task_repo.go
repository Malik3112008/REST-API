package repository

import (
	"study/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id uint) (*model.Task, error)
	UpdateStatus(id uint, status string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.Preload("Assignee").First(&task, id).Error
	return &task, err
}

func (r *taskRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.Task{}).Where("id = ?", id).Update("status", status).Error
}

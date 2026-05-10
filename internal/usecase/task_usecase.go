package usecase

import (
	"study/internal/model"
	"study/internal/repository"
)

type TaskUsecase interface {
	CreateTask(task *model.Task) error
	GetTaskByID(id uint) (*model.Task, error)
	UpdateTaskStatus(id uint, status string) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) CreateTask(task *model.Task) error {
	return u.repo.Create(task)
}

func (u *taskUsecase) GetTaskByID(id uint) (*model.Task, error) {
	return u.repo.GetByID(id)
}

func (u *taskUsecase) UpdateTaskStatus(id uint, status string) error {
	return u.repo.UpdateStatus(id, status)
}

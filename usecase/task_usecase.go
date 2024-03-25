package usecase

import (
	"context"
	"time"

	domain "github.com/ijul/be-monggo/domain/request"
)

type TaskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

// Create implements domain.TaskUsecase.
func (t *TaskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.Create(ctx, task)
}

// FetchByUserID implements domain.TaskUsecase.
func (t *TaskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.FetchByUserID(ctx, userID)
}

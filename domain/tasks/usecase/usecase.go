package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/tasks/model"
	task_repo "telegram-clicker-game-be/domain/tasks/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	taskRepo task_repo.RepoInterface
	logger   *logrus.Logger
}

type UsecaseInterface interface {
	GetTasksByUser(ctx context.Context) (res []model.TaskData, err error)
}

func NewUsecase(taskRepo task_repo.RepoInterface, logger *logrus.Logger) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(taskRepo); err != nil {
		return nil, err
	}

	return &usecase{
		taskRepo: taskRepo,
		logger:   logger,
	}, nil
}

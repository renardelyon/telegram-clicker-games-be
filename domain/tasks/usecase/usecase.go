package usecase

import (
	"context"
	"telegram-clicker-game-be/constant"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/domain/tasks/model"
	task_repo "telegram-clicker-game-be/domain/tasks/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type usecase struct {
	taskRepo     task_repo.RepoInterface
	gameplayRepo gameplay_repo.RepoInterface
	logger       *logrus.Logger
	dbClient     *mongo.Client
}

type UsecaseInterface interface {
	GetTasksByUser(ctx context.Context) (res []model.TaskData, err error)
	RedeemTaskReward(ctx context.Context, taskId string, status constant.TaskStatus) (err error)
}

func NewUsecase(taskRepo task_repo.RepoInterface, logger *logrus.Logger, gameplayRepo gameplay_repo.RepoInterface, dbClient *mongo.Client) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(taskRepo, gameplayRepo); err != nil {
		return nil, err
	}

	return &usecase{
		gameplayRepo: gameplayRepo,
		dbClient:     dbClient,
		taskRepo:     taskRepo,
		logger:       logger,
	}, nil
}

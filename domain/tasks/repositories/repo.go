package repo

import (
	"context"
	"telegram-clicker-game-be/domain/tasks/model"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	dbMongo *mongo.Database
	logger  *logrus.Logger
}

type RepoInterface interface {
	GetTasksByUser(ctx context.Context, userId int) (res []model.TaskData, err error)
	GetTaskMasterById(ctx context.Context, taskId primitive.ObjectID) (taskMaster model.TaskMaster, err error)
	UpdateUserTask(ctx context.Context, userId int, task *model.Task) (err error)
	GetUserTaskById(ctx context.Context, userId int, taskId primitive.ObjectID) (task model.Task, err error)
}

func NewRepo(dbMongo *mongo.Database, logger *logrus.Logger) (RepoInterface, error) {
	if err := utils.ExpectPointer(dbMongo); err != nil {
		return nil, err
	}

	return &repo{
		dbMongo: dbMongo,
		logger:  logger,
	}, nil
}

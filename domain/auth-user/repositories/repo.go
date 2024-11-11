package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	dbMongo *mongo.Database
	logger  *logrus.Logger
}

type RepoInterface interface {
	FindDocumentByTelegrarmId(ctx context.Context, telegramId int64) (result bson.M, err error)
	GetAllUpgrades(ctx context.Context) (result []model.UpgradeMaster, err error)
	GetAllTasks(ctx context.Context) (result []model.TaskMaster, err error)
	InserUserData(ctx context.Context, user *model.Users) (err error)
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

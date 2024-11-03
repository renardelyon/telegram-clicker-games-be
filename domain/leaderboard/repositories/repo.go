package repo

import (
	"context"
	"telegram-clicker-game-be/domain/leaderboard/model"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	dbMongo *mongo.Database
	logger  *logrus.Logger
}

type RepoInterface interface {
	GetUserWithLimit(ctx context.Context, limit int, order int) (res []model.User, err error)
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

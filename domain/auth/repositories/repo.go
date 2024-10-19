package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth/model"
	"telegram-clicker-game-be/pkg/error_utils"
	"telegram-clicker-game-be/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	dbMongo *mongo.Database
}

type RepoInterface interface {
	FindDocumentByTelegrarmId(ctx context.Context, telegramId int64) (result bson.M, err error)
	GetAllUpgrades(ctx context.Context) (result []model.UpgradeMaster, err error)
	GetAllTasks(ctx context.Context) (result []model.TaskMaster, err error)
	InserUserData(ctx context.Context, user *model.Users) (err error)
}

func NewRepo(dbMongo *mongo.Database) (RepoInterface, error) {
	if err := utils.ExpectPointer(dbMongo); err != nil {
		return nil, error_utils.HandleError(err)
	}

	return &repo{
		dbMongo: dbMongo,
	}, nil
}

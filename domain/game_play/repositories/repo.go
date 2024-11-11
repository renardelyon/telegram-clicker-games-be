package repo

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
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
	GetUpgradeMasterByEffect(ctx context.Context, effect string) (upgradeMaster model.UpgradeMaster, err error)
	GetUserUpgradesByTelegramId(ctx context.Context, userId int) (upgrades []model.Upgrade, err error)
	GetUserGameState(ctx context.Context, userId int) (states model.GameState, err error)
	UpdateBalanceGameState(ctx context.Context, userId int, state model.GameState) (err error)
	DecrementBalance(ctx context.Context, userId int, value float64) (err error)
	GetUpgradeMasterById(ctx context.Context, upgradeId primitive.ObjectID) (upgradeMaster model.UpgradeMaster, err error)
	GetUserUpgradeByUpgradeId(ctx context.Context, userId int, upgradeId primitive.ObjectID) (upgrade model.Upgrade, err error)
	UpdateUserUpgradeByUpgradeId(ctx context.Context, userId int, upgrade *model.Upgrade) (err error)
	GetUpgradesByUser(ctx context.Context, userId int) (res []model.UpgradeData, err error)
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

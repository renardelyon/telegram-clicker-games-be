package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/domain/game_play/payload"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type usecase struct {
	gameplayRepo gameplay_repo.RepoInterface
	logger       *logrus.Logger
	dbClient     *mongo.Client
}

type UsecaseInterface interface {
	SubmitTaps(ctx context.Context, taps *payload.SubmitTapsPayload) error
	BuyUpgrade(ctx context.Context, upgradeId string) (err error)
	GetUpgrades(ctx context.Context) (upgrades []model.UpgradeData, err error)
}

func NewUsecase(gameplayRepo gameplay_repo.RepoInterface, logger *logrus.Logger, dbClient *mongo.Client) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(gameplayRepo); err != nil {
		return nil, err
	}

	return &usecase{
		gameplayRepo: gameplayRepo,
		logger:       logger,
		dbClient:     dbClient,
	}, nil
}

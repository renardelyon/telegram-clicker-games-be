package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/payload"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	gameplayRepo gameplay_repo.RepoInterface
	logger       *logrus.Logger
}

type UsecaseInterface interface {
	SubmitTaps(ctx context.Context, taps *payload.SubmitTapsPayload) error
}

func NewUsecase(gameplayRepo gameplay_repo.RepoInterface, logger *logrus.Logger) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(gameplayRepo); err != nil {
		return nil, err
	}

	return &usecase{
		gameplayRepo: gameplayRepo,
		logger:       logger,
	}, nil
}

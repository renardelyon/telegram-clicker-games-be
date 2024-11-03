package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/leaderboard/model"
	leaderboard_repo "telegram-clicker-game-be/domain/leaderboard/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	leaderboardRepo leaderboard_repo.RepoInterface
	logger          *logrus.Logger
}

type UsecaseInterface interface {
	GetLeaderboard(ctx context.Context, limit int) (res []model.User, err error)
}

func NewUsecase(leaderboardRepo leaderboard_repo.RepoInterface, logger *logrus.Logger) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(leaderboardRepo); err != nil {
		return nil, err
	}

	return &usecase{
		leaderboardRepo: leaderboardRepo,
		logger:          logger,
	}, nil
}

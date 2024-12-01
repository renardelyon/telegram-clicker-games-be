package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/auth-user/model"
	auth_repo "telegram-clicker-game-be/domain/auth-user/repositories"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type usecase struct {
	authRepo     auth_repo.RepoInterface
	gameplayRepo gameplay_repo.RepoInterface
	logger       *logrus.Logger
}

type UsecaseInterface interface {
	ValidateAndBindUserInfo(ctx context.Context, telData string) (data initdata.InitData, err error)
	GetUserById(ctx context.Context) (users model.Users, err error)
	UpsertUser(ctx context.Context) (err error)
	UpdateEnergyBasedOnTime(ctx context.Context) (err error)
}

func NewUsecase(authRepo auth_repo.RepoInterface, gameplayRepo gameplay_repo.RepoInterface, logger *logrus.Logger) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(authRepo); err != nil {
		return nil, err
	}
	if err := utils.ExpectPointer(gameplayRepo); err != nil {
		return nil, err
	}

	return &usecase{
		authRepo:     authRepo,
		logger:       logger,
		gameplayRepo: gameplayRepo,
	}, nil
}

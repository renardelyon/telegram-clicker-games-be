package usecase

import (
	"context"
	repo "telegram-clicker-game-be/domain/auth/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type usecase struct {
	authRepo repo.RepoInterface
	logger   *logrus.Logger
}

type UsecaseInterface interface {
	ValidateAndBindUserInfo(ctx context.Context, telData string) (data initdata.InitData, err error)
}

func NewUsecase(authRepo repo.RepoInterface, logger *logrus.Logger) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(authRepo); err != nil {
		return nil, err
	}

	return &usecase{
		authRepo: authRepo,
		logger:   logger,
	}, nil
}

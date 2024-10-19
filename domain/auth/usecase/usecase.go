package usecase

import (
	"context"
	repo "telegram-clicker-game-be/domain/auth/repositories"
	"telegram-clicker-game-be/pkg/error_utils"
	"telegram-clicker-game-be/pkg/utils"
)

type usecase struct {
	authRepo repo.RepoInterface
}

type UsecaseInterface interface {
	ValidateAndBindUserInfo(ctx context.Context, telData string) error
}

func NewUsecase(authRepo repo.RepoInterface) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(authRepo); err != nil {
		return nil, error_utils.HandleError(err)
	}

	return &usecase{
		authRepo: authRepo,
	}, nil
}

package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/referral/model"
	referral_repo "telegram-clicker-game-be/domain/referral/repositories"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type usecase struct {
	dbClient     *mongo.Client
	referralRepo referral_repo.RepoInterface
	logger       *logrus.Logger
}

type UsecaseInterface interface {
	GetReferrals(ctx context.Context) (res []model.User, err error)
	AddReferrals(ctx context.Context, referred_by int) (err error)
}

func NewUsecase(referralRepo referral_repo.RepoInterface, logger *logrus.Logger, dbClient *mongo.Client) (UsecaseInterface, error) {
	if err := utils.ExpectPointer(referralRepo); err != nil {
		return nil, err
	}

	return &usecase{
		dbClient:     dbClient,
		referralRepo: referralRepo,
		logger:       logger,
	}, nil
}

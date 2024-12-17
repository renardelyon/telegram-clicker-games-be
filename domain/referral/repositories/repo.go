package repo

import (
	"context"

	"telegram-clicker-game-be/domain/referral/model"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	dbMongo *mongo.Database
	logger  *logrus.Logger
}

type RepoInterface interface {
	GetReferralByUserId(ctx context.Context, userId int) (res model.Referral, err error)
	GetUserByReferralUserId(ctx context.Context, userIds ...int) (res []model.User, err error)
	AddReferredBy(ctx context.Context, userId int, referrerId int) (err error)
	AddReferral(ctx context.Context, userId int, referrerId int) (err error)
	CheckReferralExist(ctx context.Context, userId int, referrerId int) (res bool, err error)
	ResetNewUserReferred(ctx context.Context, userId int) (err error)
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

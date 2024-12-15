package repo

import (
	"context"
	"telegram-clicker-game-be/config"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/domain/auth-user/response"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	httpClient *resty.Client
	cfg        *config.Config
	dbMongo    *mongo.Database
	logger     *logrus.Logger
}

type RepoInterface interface {
	FindDocumentByTelegrarmId(ctx context.Context, telegramId int64) (result model.Users, err error)
	GetAllUpgrades(ctx context.Context) (result []model.UpgradeMaster, err error)
	GetAllTasks(ctx context.Context) (result []model.TaskMaster, err error)
	InserUserData(ctx context.Context, user *model.Users) (err error)
	GetUserById(ctx context.Context, userId int) (result model.Users, err error)
	UpsertUserData(ctx context.Context, user *model.Users) (err error)
	CheckMembershipTelegram(ctx context.Context, telegramId int) (res response.TelegramMembershipResponse, err error)
}

func NewRepo(dbMongo *mongo.Database, logger *logrus.Logger, cfg *config.Config, httpClient *resty.Client) (RepoInterface, error) {
	if err := utils.ExpectPointer(dbMongo); err != nil {
		return nil, err
	}

	return &repo{
		dbMongo:    dbMongo,
		logger:     logger,
		cfg:        cfg,
		httpClient: httpClient,
	}, nil
}

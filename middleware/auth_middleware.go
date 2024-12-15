package middleware

import (
	"telegram-clicker-game-be/config"
	"telegram-clicker-game-be/domain/auth-user/handler"
	auth_repo "telegram-clicker-game-be/domain/auth-user/repositories"
	auth_usecase "telegram-clicker-game-be/domain/auth-user/usecase"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthMiddleware(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine) error {
	// ROUTING

	authRepo, err := auth_repo.NewRepo(dbMongo, logger, &config.Config{}, &resty.Client{})
	if err != nil {
		return error_utils.HandleError(err)
	}

	gameplayRepo, err := gameplay_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := auth_usecase.NewUsecase(authRepo, gameplayRepo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := handler.NewHandler(usecase)

	r.Use(handler.ValidateAndBindUserInfo)

	return nil

}

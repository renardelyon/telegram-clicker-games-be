package middleware

import (
	"telegram-clicker-game-be/domain/auth/handler"
	auth_repo "telegram-clicker-game-be/domain/auth/repositories"
	auth_usecase "telegram-clicker-game-be/domain/auth/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthMiddleware(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine) error {
	// ROUTING

	repo, err := auth_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := auth_usecase.NewUsecase(repo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := handler.NewHandler(usecase)

	r.Use(handler.ValidateAndBindUserInfo)

	return nil

}

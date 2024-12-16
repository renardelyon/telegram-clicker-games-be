package middleware

import (
	"telegram-clicker-game-be/domain/auth-user/handler"
	auth_repo "telegram-clicker-game-be/domain/auth-user/repositories"
	auth_usecase "telegram-clicker-game-be/domain/auth-user/usecase"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthMiddleware(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine,
	authRepo auth_repo.RepoInterface,
	gameplayRepo gameplay_repo.RepoInterface,
) error {
	// ROUTING

	usecase, err := auth_usecase.NewUsecase(authRepo, gameplayRepo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := handler.NewHandler(usecase)

	r.Use(handler.ValidateAndBindUserInfo)

	return nil

}

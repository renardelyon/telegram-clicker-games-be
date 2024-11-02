package route

import (
	gameplay_handler "telegram-clicker-game-be/domain/game_play/handler"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	gameplay_usecase "telegram-clicker-game-be/domain/game_play/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupGameplayRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine,
	apiRoute *gin.RouterGroup) error {
	// ROUTING

	repo, err := gameplay_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := gameplay_usecase.NewUsecase(repo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := gameplay_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		v1.PUT("/submit-taps", handler.SubmitTaps)
	}

	return nil
}

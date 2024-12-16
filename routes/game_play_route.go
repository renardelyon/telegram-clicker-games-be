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
	dbClient *mongo.Client,
	r *gin.Engine,
	apiRoute *gin.RouterGroup,
	gameplayRepo gameplay_repo.RepoInterface,
) error {
	// ROUTING

	usecase, err := gameplay_usecase.NewUsecase(gameplayRepo, logger, dbClient)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := gameplay_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		gameplay := v1.Group("/gameplay")
		gameplay.PUT("/submit-taps", handler.SubmitTaps)
		gameplay.PUT("/buy-upgrade", handler.BuyUpgrade)
		gameplay.GET("/upgrades", handler.GetTasksByUser)
	}

	return nil
}

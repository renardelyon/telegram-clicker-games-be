package route

import (
	leaderboard_handler "telegram-clicker-game-be/domain/leaderboard/handler"
	leaderboard_repo "telegram-clicker-game-be/domain/leaderboard/repositories"
	leaderboard_usecase "telegram-clicker-game-be/domain/leaderboard/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupLeaderboardRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine,
	apiRoute *gin.RouterGroup) error {
	// ROUTING

	repo, err := leaderboard_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := leaderboard_usecase.NewUsecase(repo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := leaderboard_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		v1.GET("/leaderboard", handler.GetLeaderboard)
	}

	return nil
}
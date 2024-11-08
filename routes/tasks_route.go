package route

import (
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	tasks_handler "telegram-clicker-game-be/domain/tasks/handler"
	tasks_repo "telegram-clicker-game-be/domain/tasks/repositories"
	tasks_usecase "telegram-clicker-game-be/domain/tasks/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupTasksRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	dbClient *mongo.Client,
	r *gin.Engine,
	apiRoute *gin.RouterGroup) error {
	// ROUTING

	repo, err := tasks_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	gRepo, err := gameplay_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := tasks_usecase.NewUsecase(repo, logger, gRepo, dbClient)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := tasks_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		v1.GET("/tasks", handler.GetTasksByUser)
		v1.PUT("/redeem-task", handler.RedeemTaskReward)
	}

	return nil
}

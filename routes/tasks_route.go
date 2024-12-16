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
	apiRoute *gin.RouterGroup,
	taskRepo tasks_repo.RepoInterface,
	gameplayRepo gameplay_repo.RepoInterface,
) error {
	// ROUTING

	usecase, err := tasks_usecase.NewUsecase(taskRepo, logger, gameplayRepo, dbClient)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := tasks_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		taskRoute := v1.Group("/tasks")
		taskRoute.GET("/list", handler.GetTasksByUser)
		taskRoute.PUT("/redeem", handler.RedeemTaskReward)
	}

	return nil
}

package route

import (
	user_auth_handler "telegram-clicker-game-be/domain/auth-user/handler"
	user_auth_repo "telegram-clicker-game-be/domain/auth-user/repositories"
	user_auth_usecase "telegram-clicker-game-be/domain/auth-user/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	dbClient *mongo.Client,
	r *gin.Engine,
	apiRoute *gin.RouterGroup) error {
	// ROUTING

	repo, err := user_auth_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := user_auth_usecase.NewUsecase(repo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := user_auth_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		v1.GET("/profile", handler.GetUserById)
		v1.POST("/sign-in", handler.SignIn)
	}

	return nil
}

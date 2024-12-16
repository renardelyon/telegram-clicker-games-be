package route

import (
	"telegram-clicker-game-be/config"
	user_auth_handler "telegram-clicker-game-be/domain/auth-user/handler"
	user_auth_repo "telegram-clicker-game-be/domain/auth-user/repositories"
	user_auth_usecase "telegram-clicker-game-be/domain/auth-user/usecase"
	gameplay_repo "telegram-clicker-game-be/domain/game_play/repositories"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	r *gin.Engine,
	apiRoute *gin.RouterGroup,
	cfg *config.Config,
	authRepo user_auth_repo.RepoInterface,
	gameplayRepo gameplay_repo.RepoInterface,
) error {
	// ROUTING

	usecase, err := user_auth_usecase.NewUsecase(authRepo, gameplayRepo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := user_auth_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		authGroup.GET("/profile", handler.UpdateEnergyBasedOnTime, handler.GetUserById)
		authGroup.POST("/sign-in", handler.SignIn)
		authGroup.GET("/telegram/check", handler.CheckMembershipTelegram)
	}

	return nil
}

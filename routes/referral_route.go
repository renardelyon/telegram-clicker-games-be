package route

import (
	referral_handler "telegram-clicker-game-be/domain/referral/handler"
	referral_repo "telegram-clicker-game-be/domain/referral/repositories"
	referral_usecase "telegram-clicker-game-be/domain/referral/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupReferralRoute(
	logger *logrus.Logger,
	dbMongo *mongo.Database,
	dbClient *mongo.Client,
	r *gin.Engine,
	apiRoute *gin.RouterGroup) error {
	// ROUTING

	repo, err := referral_repo.NewRepo(dbMongo, logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	usecase, err := referral_usecase.NewUsecase(repo, logger, dbClient)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := referral_handler.NewHandler(usecase)

	v1 := apiRoute.Group("/v1")
	{
		v1.GET("/referral", handler.GetReferrals)
		v1.POST("/referral", handler.AddReferrals)
	}

	return nil
}

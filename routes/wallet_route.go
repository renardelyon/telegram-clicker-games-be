package route

import (
	"telegram-clicker-game-be/config"
	wallet_handler "telegram-clicker-game-be/domain/wallets/handler"
	"telegram-clicker-game-be/domain/wallets/usecase"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupWalletRoute(
	cfg *config.Config,
	logger *logrus.Logger,
	r *gin.Engine,
) error {
	// ROUTING

	usecase, err := usecase.NewUsecase(logger)
	if err != nil {
		return error_utils.HandleError(err)
	}

	handler := wallet_handler.NewHandler(cfg, usecase)

	group := r.Group("/wallets")
	{
		group.GET("/onConnect", handler.RedirectWalletOnConnect)
		group.GET("/onDisconnect", handler.RedirectWalletOnDisconnect)
		group.GET("/signTransaction", handler.RedirectWalletSignTransaction)
	}

	return nil
}

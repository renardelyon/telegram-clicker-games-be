package handler

import (
	"telegram-clicker-game-be/config"
	"telegram-clicker-game-be/domain/wallets/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg     *config.Config
	usecase usecase.UsecaseInterface
}

type Handler interface {
	RedirectWalletOnConnect(c *gin.Context)
	RedirectWalletOnDisconnect(c *gin.Context)
	RedirectWalletSignTransaction(c *gin.Context)
}

func NewHandler(cfg *config.Config, usecase usecase.UsecaseInterface) Handler {
	return &handler{
		usecase: usecase,
		cfg:     cfg,
	}
}

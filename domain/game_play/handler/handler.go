package handler

import (
	"telegram-clicker-game-be/domain/game_play/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	u usecase.UsecaseInterface
}

type Handler interface {
	SubmitTaps(c *gin.Context)
	BuyUpgrade(c *gin.Context)
}

func NewHandler(usecase usecase.UsecaseInterface) Handler {
	return &handler{
		u: usecase,
	}
}

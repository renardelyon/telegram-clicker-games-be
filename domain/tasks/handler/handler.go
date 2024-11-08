package handler

import (
	"telegram-clicker-game-be/domain/tasks/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	u usecase.UsecaseInterface
}

type Handler interface {
	GetTasksByUser(c *gin.Context)
	RedeemTaskReward(c *gin.Context)
}

func NewHandler(usecase usecase.UsecaseInterface) Handler {
	return &handler{
		u: usecase,
	}
}

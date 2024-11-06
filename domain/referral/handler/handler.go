package handler

import (
	"telegram-clicker-game-be/domain/referral/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	u usecase.UsecaseInterface
}

type Handler interface {
	GetReferrals(c *gin.Context)
	AddReferrals(c *gin.Context)
}

func NewHandler(usecase usecase.UsecaseInterface) Handler {
	return &handler{
		u: usecase,
	}
}

package handler

import (
	"telegram-clicker-game-be/domain/auth-user/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	u usecase.UsecaseInterface
}

type Handler interface {
	ValidateAndBindUserInfo(c *gin.Context)
	GetUserById(c *gin.Context)
	SignIn(c *gin.Context)
}

func NewHandler(usecase usecase.UsecaseInterface) Handler {
	return &handler{
		u: usecase,
	}
}

package handler

import (
	"net/http"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetUserById(c *gin.Context) {
	user, err := h.u.GetUserById(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			utils.NewResponse(utils.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}),
		)
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(utils.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	}))
}

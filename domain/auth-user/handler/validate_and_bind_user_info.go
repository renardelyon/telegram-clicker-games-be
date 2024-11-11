package handler

import (
	"net/http"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) ValidateAndBindUserInfo(c *gin.Context) {
	telData := c.GetHeader("X-init-telegram-data")

	data, err := h.u.ValidateAndBindUserInfo(c, telData)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			utils.NewResponse(utils.Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}),
		)
		return
	}

	c.Set("user_info", &data)

	c.Next()
}

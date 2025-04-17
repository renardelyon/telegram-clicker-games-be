package handler

import (
	"fmt"
	"net/http"
	"telegram-clicker-game-be/domain/wallets/params"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) RedirectWalletOnConnect(c *gin.Context) {
	var params params.Params

	err := c.ShouldBindQuery(&params)
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

	encodedqp, err := h.usecase.EncodQueryParams(c, params)
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

	c.Redirect(http.StatusFound, fmt.Sprintf("%s?startapp=%s", h.cfg.Telegram.BotApp, encodedqp))
}

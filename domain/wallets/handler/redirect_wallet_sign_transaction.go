package handler

import (
	"fmt"
	"net/http"
	"telegram-clicker-game-be/domain/wallets/params"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) RedirectWalletSignTransaction(c *gin.Context) {
	var walletErr params.WalletError

	err := c.ShouldBindQuery(&walletErr)
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

	if walletErr.ErrorCode != nil {
		encodedqp, err := h.usecase.EncodQueryParams(c, walletErr)
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
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("%s?startapp=%s", h.cfg.Telegram.BotApp, "transaction_completed"))
}

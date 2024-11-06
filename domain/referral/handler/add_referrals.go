package handler

import (
	"net/http"
	"telegram-clicker-game-be/domain/referral/payload"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddReferrals(c *gin.Context) {
	var payload payload.AddReferralsPayload
	err := c.ShouldBindJSON(&payload)
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

	err = h.u.AddReferrals(c, payload.ReferredBy)
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
	}))
}

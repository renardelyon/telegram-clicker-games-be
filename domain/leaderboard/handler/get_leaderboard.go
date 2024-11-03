package handler

import (
	"net/http"
	"telegram-clicker-game-be/domain/leaderboard/payload"
	"telegram-clicker-game-be/domain/leaderboard/response"
	"telegram-clicker-game-be/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetLeaderboard(c *gin.Context) {
	var payload payload.GetLeaderboardPayload
	err := c.ShouldBindQuery(&payload)
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

	users, err := h.u.GetLeaderboard(c, payload.Limit)
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

	c.JSON(http.StatusOK, utils.NewResponse(utils.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response.GetLeaderboardResponse{Users: users},
	}))
}

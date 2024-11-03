package response

import "telegram-clicker-game-be/domain/leaderboard/model"

type GetLeaderboardResponse struct {
	Users []model.User `json:"users"`
}

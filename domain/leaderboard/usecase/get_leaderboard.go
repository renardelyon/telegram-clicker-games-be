package usecase

import (
	"context"
	"telegram-clicker-game-be/constant"
	"telegram-clicker-game-be/domain/leaderboard/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
)

func (u *usecase) GetLeaderboard(ctx context.Context, limit int) (res []model.User, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"limit":      limit,
	}).Info("Usecase: GetLeaderboard")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, u.logger)

	// TODO: Caching

	res, err = u.leaderboardRepo.GetUserWithLimit(ctx, limit, constant.DESC)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

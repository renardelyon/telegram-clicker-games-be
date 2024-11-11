package usecase

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) GetUpgrades(ctx context.Context) (upgrades []model.UpgradeData, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
	}).Info("Usecase: GetUpgrades")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	upgrades, err = u.gameplayRepo.GetUpgradesByUser(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

package usecase

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/constant"
	gameplay_model "telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/domain/game_play/payload"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) SubmitTaps(ctx context.Context, taps *payload.SubmitTapsPayload) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"data":       fmt.Sprintf("%+v", taps),
	}).Info("Usecase: SubmitTaps")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	upgradeMaster, err := u.gameplayRepo.GetUpgradeMasterByEffect(ctx, constant.MULTI_TAP)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	upgrades, err := u.gameplayRepo.GetUserUpgradesByTelegramId(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	var multiTapUpgrade gameplay_model.Upgrade
	for _, up := range upgrades {
		if up.UpgradeId == upgradeMaster.Id {
			multiTapUpgrade = up
			break
		}
	}

	gameStates, err := u.gameplayRepo.GetUserGameState(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// update game states struct
	gameStates.Balance += float64(multiTapUpgrade.Level) * float64(taps.Taps)
	gameStates.ClickCount += int64(taps.Taps)
	gameStates.TotalBalance += float64(multiTapUpgrade.Level) * float64(taps.Taps)

	err = u.gameplayRepo.UpdateBalanceGameState(ctx, int(userInfo.ID), gameStates)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return nil
}

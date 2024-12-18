package usecase

import (
	"context"
	"errors"
	"fmt"
	"telegram-clicker-game-be/constant"
	gameplay_model "telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/domain/game_play/payload"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) SubmitTaps(ctx context.Context, taps *payload.SubmitTapsPayload) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"data":       fmt.Sprintf("%+v", taps),
	}).Info("Usecase: SubmitTaps")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	multiTapUpgradeMaster, err := u.gameplayRepo.GetUpgradeMasterByEffect(ctx, constant.MULTI_TAP)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	energyRechargeUpgradeMaster, err := u.gameplayRepo.GetUpgradeMasterByEffect(ctx, constant.ENERGY_RECHARGE)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	energyLimitUpgradeMaster, err := u.gameplayRepo.GetUpgradeMasterByEffect(ctx, constant.ENERGY_LIMIT)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	upgrades, err := u.gameplayRepo.GetUserUpgradesByTelegramId(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	var (
		multiTapUpgrade       gameplay_model.Upgrade
		energyRechargeUpgrade gameplay_model.Upgrade
		energyLimitUpgrade    gameplay_model.Upgrade
	)

	for _, up := range upgrades {
		switch up.UpgradeId {
		case multiTapUpgradeMaster.Id:
			multiTapUpgrade = up
			break
		case energyRechargeUpgradeMaster.Id:
			energyRechargeUpgrade = up
			break
		case energyLimitUpgradeMaster.Id:
			energyLimitUpgrade = up
			break
		default:
			break
		}
	}

	gameStates, err := u.gameplayRepo.GetUserGameState(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// Adjust energy based on recharge and consumption
	updatedEnergy, err := u.calculateEnergy(&gameStates, taps, &multiTapUpgrade, &energyRechargeUpgrade, &energyLimitUpgrade)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// update game states struct
	gameStates.Balance += float64(multiTapUpgrade.Level) * float64(taps.Taps)
	gameStates.ClickCount += int64(taps.Taps)
	gameStates.TotalBalance += float64(multiTapUpgrade.Level) * float64(taps.Taps)
	gameStates.Energy = updatedEnergy
	gameStates.LastEnergyUpdate = time.Now()

	err = u.gameplayRepo.UpdateBalanceGameState(ctx, int(userInfo.ID), gameStates)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return nil
}

func (u *usecase) calculateEnergy(
	gameStates *gameplay_model.GameState,
	taps *payload.SubmitTapsPayload,
	multiTapUpgrade *gameplay_model.Upgrade,
	energyRechargeUpgrade *gameplay_model.Upgrade,
	energyLimitUpgrade *gameplay_model.Upgrade,
) (res int32, err error) {
	energyDiff := int32(taps.Time.ConvertToGoTime().Sub(gameStates.LastEnergyUpdate).Seconds()) *
		energyRechargeUpgrade.Level

	newEnergy := func() int32 {
		maxEnergy := float64(gameStates.BaseEnergy) * float64(energyLimitUpgrade.Level)
		energyTotal := gameStates.Energy + int32(energyDiff)
		if energyTotal > int32(maxEnergy) {
			return int32(maxEnergy)
		}

		return energyTotal
	}()

	res = newEnergy - int32(taps.Taps)*multiTapUpgrade.Level

	if res < 0 {
		return res, errors.New("energy is less than zero, cannot increment taps anymore")
	}

	return
}

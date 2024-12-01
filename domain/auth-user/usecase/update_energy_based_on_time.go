package usecase

import (
	"context"
	"telegram-clicker-game-be/constant"
	gameplay_model "telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) UpdateEnergyBasedOnTime(ctx context.Context) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
	}).Info("Usecase: GetUserById")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

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
		energyRechargeUpgrade gameplay_model.Upgrade
		energyLimitUpgrade    gameplay_model.Upgrade
	)

	for _, up := range upgrades {
		switch up.UpgradeId {
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

	newEnergy := u.calculateEnergyBasedOnNow(&gameStates, &energyRechargeUpgrade, &energyLimitUpgrade)

	gameStates.Energy = newEnergy
	gameStates.LastEnergyUpdate = time.Now()

	err = u.gameplayRepo.UpdateBalanceGameState(ctx, int(userInfo.ID), gameStates)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

func (u *usecase) calculateEnergyBasedOnNow(
	gameStates *gameplay_model.GameState,
	energyRechargeUpgrade *gameplay_model.Upgrade,
	energyLimitUpgrade *gameplay_model.Upgrade,
) (res int32) {
	energyDiff := int32(time.Now().Sub(gameStates.LastEnergyUpdate).Seconds()) *
		energyRechargeUpgrade.Level

	newEnergy := func() int32 {
		maxEnergy := float64(gameStates.BaseEnergy) * float64(energyLimitUpgrade.Level)
		energyTotal := gameStates.Energy + int32(energyDiff)
		if energyTotal > int32(maxEnergy) {
			return int32(maxEnergy)
		}

		return energyTotal
	}()

	return newEnergy
}

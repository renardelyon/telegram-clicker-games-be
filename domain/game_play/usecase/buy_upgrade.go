package usecase

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (u *usecase) BuyUpgrade(ctx context.Context, upgradeId string) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"upgradeId":  upgradeId,
	}).Info("Usecase: BuyUpgrade")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, u.logger)

	now := time.Now()

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	uId, err := primitive.ObjectIDFromHex(upgradeId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// Get Upgrades
	upgrade, err := u.gameplayRepo.GetUserUpgradeByUpgradeId(ctx, int(userInfo.ID), uId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// Get UpgradeMaster
	upgradeMaster, err := u.gameplayRepo.GetUpgradeMasterById(ctx, uId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	/**
	    - Change Next Cost to NC = BaseCost * inc_multiplier * Level
	    - Increment Level
	    - Update acquired_at
	**/
	upgrade.AcquiredAt = now
	upgrade.NextCost = upgradeMaster.BaseCost * upgradeMaster.IncMultiplier * float64(upgrade.Level)
	upgrade.Level += 1

	operation := func(mctx mongo.SessionContext) (res interface{}, err error) {
		// Decrement Balance
		// If Balance Less than 0 throw error
		err = u.gameplayRepo.DecrementBalance(ctx, int(userInfo.ID), upgrade.NextCost)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return
		}

		err = u.gameplayRepo.UpdateUserUpgradeByUpgradeId(ctx, int(userInfo.ID), &upgrade)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return
		}

		return
	}

	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)

	// Starts a session on the client
	session, err := u.dbClient.StartSession()
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}
	// Defers ending the session after the transaction is committed or ended
	defer session.EndSession(ctx)

	// Inserts multiple documents into a collection within a transaction,
	// then commits or ends the transaction
	_, err = session.WithTransaction(ctx, operation, txnOptions)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

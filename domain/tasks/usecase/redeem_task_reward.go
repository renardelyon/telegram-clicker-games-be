package usecase

import (
	"context"
	"errors"
	"telegram-clicker-game-be/constant"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (u *usecase) RedeemTaskReward(ctx context.Context, taskId string, status constant.TaskStatus) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"task_id":    taskId,
	}).Info("Usecase: RedeemTaskReward")

	now := time.Now()

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	tId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// Get Task Master
	tm, err := u.taskRepo.GetTaskMasterById(ctx, tId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// Get User Task
	ut, err := u.taskRepo.GetUserTaskById(ctx, int(userInfo.ID), tId)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	if !status.IsValid() {
		err = errors.New("task status doesn't exist")
		errTrace = error_utils.HandleError(err)
		return
	}

	ut.Status = string(status)
	ut.LastUpdated = now

	// Get User Game State
	gs, err := u.gameplayRepo.GetUserGameState(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	operation := func(mctx mongo.SessionContext) (res interface{}, err error) {
		taskDesc := constant.TaskDesc(tm.Description)

		var reward = tm.RewardAmount
		if taskDesc.IsValid() && taskDesc == constant.INVITE_FRIENDS {
			referral, err := u.referralRepo.GetReferralByUserId(mctx, int(userInfo.ID))
			if err != nil {
				errTrace = error_utils.HandleError(err)
				return res, err
			}

			reward = tm.RewardAmount * float32(referral.NewUserReferred)

			err = u.referralRepo.ResetNewUserReferred(mctx, int(userInfo.ID))
			if err != nil {
				errTrace = error_utils.HandleError(err)
				return res, err
			}
		}

		gs.Balance += float64(reward)
		gs.TotalBalance += float64(reward)
		// Update Balance and total balance from from reward amount
		err = u.gameplayRepo.UpdateBalanceGameState(mctx, int(userInfo.ID), gs)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return
		}

		// Update task status and last_updated
		err = u.taskRepo.UpdateUserTask(mctx, int(userInfo.ID), &ut)
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

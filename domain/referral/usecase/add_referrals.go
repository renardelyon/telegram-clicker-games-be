package usecase

import (
	"context"

	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (u *usecase) AddReferrals(ctx context.Context, referred_by int) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id":  ctx.Value("request_id"),
		"referred_by": referred_by,
	}).Info("Usecase: AddReferrals")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

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

	operation := func(mctx mongo.SessionContext) (res interface{}, err error) {
		if err = u.referralRepo.AddReferredBy(mctx, int(userInfo.ID), referred_by); err != nil {
			return
		}

		if err = u.referralRepo.AddReferral(mctx, int(userInfo.ID), referred_by); err != nil {
			return
		}

		return
	}

	// Inserts multiple documents into a collection within a transaction,
	// then commits or ends the transaction
	_, err = session.WithTransaction(ctx, operation, txnOptions)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

package repo

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) AddReferredBy(ctx context.Context, userId int, referrerId int) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"referrerId": referrerId,
	}).Info("Repo: AddReferredBy")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	now := time.Now()

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": userId}

	// Update data
	update := bson.M{
		"$set": bson.M{
			"referral.referred_by": referrerId,
			"updated_at":           now,
		},
	}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

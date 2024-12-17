package repo

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) ResetNewUserReferred(ctx context.Context, userId int) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
	}).Info("Repo: ResetNewUserReferred")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	now := time.Now()

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": userId}

	// Update data
	update := bson.M{
		"$set": bson.M{
			"updated_at":        now,
			"new_user_referred": 0,
		},
	}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

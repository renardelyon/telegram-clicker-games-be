package repo

import (
	"context"
	"telegram-clicker-game-be/domain/referral/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetUserByReferralUserId(ctx context.Context, userIds ...int) (res []model.User, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id":        ctx.Value("request_id"),
		"user_telegram_ids": userIds,
	}).Info("Repo: GetUserByReferralUserId")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": bson.M{"$in": userIds}}

	projection := options.Find().SetProjection(bson.M{
		"_id":         1,
		"telegram_id": 1,
		"first_name":  1,
		"last_name":   1,
		"user_name":   1,
	})

	cursor, err := coll.Find(ctx, filter, projection)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &res)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

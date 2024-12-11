package repo

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) CheckReferralExist(ctx context.Context, userId int, referrerId int) (res bool, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"referrerId": referrerId,
	}).Info("Repo: CheckReferral")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Define the filter
	filter := bson.M{
		"telegram_id": referrerId,
		"referral.referrals": bson.M{
			"$elemMatch": bson.M{
				"$eq": userId,
			},
		},
	}

	// Query the database
	var result bson.M
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return res, nil // No match found
		}
		errTrace = error_utils.HandleError(err)
		return
	}

	res = true
	return
}

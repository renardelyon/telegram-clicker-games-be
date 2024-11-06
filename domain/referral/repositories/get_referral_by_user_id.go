package repo

import (
	"context"
	"telegram-clicker-game-be/domain/referral/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetReferralByUserId(ctx context.Context, userId int) (res model.Referral, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id":  ctx.Value("request_id"),
		"telegram_id": userId,
	}).Info("Repo: GetReferralByUserId")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": userId}

	projection := options.FindOne().SetProjection(bson.M{"referral": 1, "_id": 0})

	var data struct {
		Referral model.Referral `bson:"referral"`
	}
	err = coll.FindOne(ctx, filter, projection).Decode(&data)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	res = data.Referral

	return
}

package repo

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetUserUpgradeByUpgradeId(ctx context.Context, userId int, upgradeId primitive.ObjectID) (upgrade model.Upgrade, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"upgradeId":  upgradeId,
	}).Info("Repo: GetUserUpgradeByUserId")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Define the filter to match the user by ID
	filter := bson.M{"telegram_id": userId}

	projection := options.
		FindOne().
		SetProjection(
			bson.M{
				"upgrades": bson.M{
					"$elemMatch": bson.M{"upgrade_id": upgradeId},
				},
			},
		)

	var data struct {
		Upgrade []model.Upgrade `bson:"upgrades"`
	}

	err = coll.FindOne(ctx, filter, projection).Decode(&data)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	if len(data.Upgrade) > 0 {
		upgrade = data.Upgrade[0]
	}

	return
}

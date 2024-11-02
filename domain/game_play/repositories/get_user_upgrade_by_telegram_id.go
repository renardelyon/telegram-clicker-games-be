package repo

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetUserUpgradesByTelegramId(ctx context.Context, userId int) (upgrades []model.Upgrade, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
	}).Info("Repo: GetUserUpgradesByTelegramId")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": userId}

	projection := options.FindOne().SetProjection(bson.M{"upgrades": 1, "_id": 0})

	var data struct {
		Upgrades []model.Upgrade `bson:"upgrades"`
	}
	// var dataBson bson.M
	err = coll.FindOne(ctx, filter, projection).Decode(&data)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	upgrades = data.Upgrades

	return
}

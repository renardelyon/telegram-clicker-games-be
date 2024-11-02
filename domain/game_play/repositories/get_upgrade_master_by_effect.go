package repo

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) GetUpgradeMasterByEffect(ctx context.Context, effect string) (upgradeMaster model.UpgradeMaster, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"effect":     effect,
	}).Info("Repo: GetUpgradeByEffect")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Upgrades")

	filter := bson.M{
		"effect": primitive.Regex{
			Pattern: effect,
			Options: "i",
		},
	}

	var upgrade model.UpgradeMaster
	err = coll.FindOne(ctx, filter).Decode(&upgrade)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return upgrade, nil
}

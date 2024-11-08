package repo

import (
	"context"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) GetUpgradeMasterById(ctx context.Context, upgradeId primitive.ObjectID) (upgradeMaster model.UpgradeMaster, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"upgradeId":  upgradeId,
	}).Info("Repo: GetUpgradeMasterById")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Upgrades")

	filter := bson.M{"_id": upgradeId}

	var upgrade model.UpgradeMaster
	err = coll.FindOne(ctx, filter).Decode(&upgrade)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return upgrade, nil
}

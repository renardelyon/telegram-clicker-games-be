package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) UpdateUserUpgradeByUpgradeId(
	ctx context.Context,
	userId int,
	upgrade *model.Upgrade,
) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"upgrade":    fmt.Sprintf("%+v", upgrade),
	}).Info("Repo: UpdateUserUpgradeByUpgradeId")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Filter to match the user by ID
	filter := bson.M{"telegram_id": userId}

	// Define the filter to check if the balance is sufficient
	// Define the update using array filters
	update := bson.M{
		"$set": bson.M{
			"upgrades.$[elem].level":       upgrade.Level,
			"upgrades.$[elem].next_cost":   upgrade.NextCost,
			"upgrades.$[elem].acquired_at": upgrade.AcquiredAt,
		},
	}

	// Specify array filters to match the upgrade by upgrade_id
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.upgrade_id": upgrade.UpgradeId}},
	}

	// Create the update options with array filters
	updateOptions := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	// Perform the update operation
	_, err = coll.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/auth/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAllUpgrades(ctx context.Context) (result []model.UpgradeMaster, err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"result":     fmt.Sprintf("%+v", result),
			"request_id": ctx.Value("request_id"),
		}).
		Info("Repo: GetAllUpgrades")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)

	coll := r.dbMongo.Collection("Upgrades")

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		errorTrace = error_utils.HandleError(err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var upgrade model.UpgradeMaster
		if err = cursor.Decode(&upgrade); err != nil {
			errorTrace = error_utils.HandleError(err)
			return
		}

		result = append(result, upgrade)
	}

	return
}

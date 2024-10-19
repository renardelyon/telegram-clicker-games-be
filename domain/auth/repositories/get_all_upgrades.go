package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAllUpgrades(ctx context.Context) (result []model.UpgradeMaster, err error) {
	coll := r.dbMongo.Collection("Upgrades")

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var upgrade model.UpgradeMaster
		if err = cursor.Decode(&upgrade); err != nil {
			return
		}

		result = append(result, upgrade)
	}

	return
}

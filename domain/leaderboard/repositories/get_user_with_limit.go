package repo

import (
	"context"
	"telegram-clicker-game-be/domain/leaderboard/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetUserWithLimit(ctx context.Context, limit int, order int) (res []model.User, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"limit":      limit,
		"order":      order,
	}).Info("Repo: GetUserWithLimit")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	options := options.Find().
		SetSort(bson.D{primitive.E{Key: "game_states.total_balance", Value: order}}).
		SetLimit(int64(limit))

	cursor, err := coll.Find(ctx, bson.D{}, options)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User

		err = cursor.Decode(&user)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return
		}

		res = append(res, user)
	}

	// Check for any errors in the cursor
	if err = cursor.Err(); err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

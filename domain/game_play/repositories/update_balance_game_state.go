package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/game_play/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) UpdateBalanceGameState(ctx context.Context, userId int, state model.GameState) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"state":      fmt.Sprintf("%+v", state),
	}).Info("Repo: UpdateBalanceGameState")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	update := bson.M{
		"$set": bson.M{"game_states": state},
	}

	filter := bson.M{"telegram_id": userId}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

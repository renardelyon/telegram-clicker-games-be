package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) DecrementBalance(ctx context.Context, userId int, value float64) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"value":      value,
	}).Info("Repo: DecrementBalance")

	var errTrace error
	defer error_utils.HandleErrorLog(errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Define the filter to check if the balance is sufficient
	filter := bson.M{
		"telegram_id":         userId,
		"game_states.balance": bson.M{"$gte": value},
	}

	// Define the update to decrement the balance
	update := bson.M{
		"$inc": bson.M{"game_states.balance": -value},
	}

	// Perform the update operation
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("insufficient balance")
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

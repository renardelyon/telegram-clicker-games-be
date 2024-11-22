package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) UpsertUserData(ctx context.Context, user *model.Users) (err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"user":       fmt.Sprintf("%+v", user),
			"request_id": ctx.Value("request_id"),
		}).
		Info("Repo: UpsertUserData")
	coll := r.dbMongo.Collection("Users")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"telegram_id": user.TelegramId}

	_, err = coll.UpdateOne(ctx, filter, bson.M{"$set": user}, opts)
	if err != nil {
		errorTrace = error_utils.HandleError(err)
		return err
	}

	return
}

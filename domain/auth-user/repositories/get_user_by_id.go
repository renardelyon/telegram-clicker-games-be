package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetUserById(ctx context.Context, userId int) (result model.Users, err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"telegram_id": userId,
			"request_id":  ctx.Value("request_id"),
		}).
		Info("Repo: GetUserById")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	filter := bson.M{"telegram_id": userId}

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		errorTrace = error_utils.HandleError(err)
		return
	}

	return
}

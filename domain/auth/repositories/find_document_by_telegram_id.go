package repo

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) FindDocumentByTelegrarmId(ctx context.Context, telegramId int64) (result bson.M, err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"telegram_id": telegramId,
			"request_id":  ctx.Value("request_id"),
		}).
		Info("Repo: FindDocumentByTelegrarmId")

	var errorTrace error
	defer error_utils.HandleErrorLog(errorTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	err = coll.FindOne(ctx, bson.M{"telegram_id": telegramId}).Decode(&result)
	if err != nil {
		errorTrace = error_utils.HandleError(err)
	}
	return
}

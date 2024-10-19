package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) FindDocumentByTelegrarmId(ctx context.Context, telegramId int64) (result bson.M, err error) {
	coll := r.dbMongo.Collection("Users")

	err = coll.FindOne(ctx, bson.M{"telegram_id": telegramId}).Decode(&result)
	return
}

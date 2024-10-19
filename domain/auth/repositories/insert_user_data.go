package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth/model"
)

func (r *repo) InserUserData(ctx context.Context, user *model.Users) (err error) {
	coll := r.dbMongo.Collection("Users")

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return
}

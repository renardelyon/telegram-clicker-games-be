package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
)

func (r *repo) InserUserData(ctx context.Context, user *model.Users) (err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"user":       fmt.Sprintf("%+v", user),
			"request_id": ctx.Value("request_id"),
		}).
		Info("Repo: GetAllUpgrades")
	coll := r.dbMongo.Collection("Users")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		errorTrace = error_utils.HandleError(err)
		return err
	}

	return
}

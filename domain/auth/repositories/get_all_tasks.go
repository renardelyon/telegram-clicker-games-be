package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/auth/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAllTasks(ctx context.Context) (result []model.TaskMaster, err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"request_id": ctx.Value("request_id"),
			"result":     fmt.Sprintf("%+v", result),
		}).
		Info("Repo: GetAllTasks")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)

	coll := r.dbMongo.Collection("Tasks")

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		errorTrace = error_utils.HandleError(err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task model.TaskMaster
		if err = cursor.Decode(&task); err != nil {
			errorTrace = error_utils.HandleError(err)
			return
		}

		result = append(result, task)
	}

	return
}

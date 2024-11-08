package repo

import (
	"context"
	"telegram-clicker-game-be/domain/tasks/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetUserTaskById(ctx context.Context, userId int, taskId primitive.ObjectID) (task model.Task, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"task_id":    taskId,
	}).Info("Repo: GetUserTaskById")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Define the filter to match the user by ID
	filter := bson.M{"telegram_id": userId}

	projection := options.
		FindOne().
		SetProjection(
			bson.M{
				"tasks": bson.M{
					"$elemMatch": bson.M{"task_id": taskId},
				},
			},
		)

	var data struct {
		Tasks []model.Task `bson:"tasks"`
	}

	err = coll.FindOne(ctx, filter, projection).Decode(&data)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	if len(data.Tasks) > 0 {
		task = data.Tasks[0]
	}

	return
}

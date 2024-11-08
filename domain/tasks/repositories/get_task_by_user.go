package repo

import (
	"context"
	"telegram-clicker-game-be/domain/tasks/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) GetTaskByUser(ctx context.Context, userId int) (res []model.TaskData, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id":  ctx.Value("request_id"),
		"telegram_id": userId,
	}).Info("Repo: GetTaskByUser")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"telegram_id": userId}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$tasks"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Tasks"},
			{Key: "localField", Value: "tasks.task_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "task_detail"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$task_detail"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &res); err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

package repo

import (
	"context"
	"telegram-clicker-game-be/domain/tasks/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) GetTaskMasterById(ctx context.Context, taskId primitive.ObjectID) (taskMaster model.TaskMaster, err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"task_id":    taskId,
	}).Info("Repo: GetTaskMasterById")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Tasks")

	filter := bson.M{"_id": taskId}

	err = coll.FindOne(ctx, filter).Decode(&taskMaster)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

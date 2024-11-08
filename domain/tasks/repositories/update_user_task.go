package repo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/domain/tasks/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) UpdateUserTask(
	ctx context.Context,
	userId int,
	task *model.Task,
) (err error) {
	r.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"userId":     userId,
		"task":       fmt.Sprintf("%+v", task),
	}).Info("Repo: UpdateUserTask")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, r.logger)

	coll := r.dbMongo.Collection("Users")

	// Filter to match the user by ID
	filter := bson.M{"telegram_id": userId}

	// Define the filter to check if the balance is sufficient
	// Define the update using array filters
	update := bson.M{
		"$set": bson.M{
			"tasks.$[elem].last_updated": task.LastUpdated,
			"tasks.$[elem].status":       task.Status,
		},
	}

	// Specify array filters to match the upgrade by upgrade_id
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.task_id": task.TaskId}},
	}

	// Create the update options with array filters
	updateOptions := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	// Perform the update operation
	_, err = coll.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}

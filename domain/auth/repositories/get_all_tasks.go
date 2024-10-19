package repo

import (
	"context"
	"telegram-clicker-game-be/domain/auth/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAllTasks(ctx context.Context) (result []model.TaskMaster, err error) {
	coll := r.dbMongo.Collection("Tasks")

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task model.TaskMaster
		if err = cursor.Decode(&task); err != nil {
			return
		}

		result = append(result, task)
	}

	return
}

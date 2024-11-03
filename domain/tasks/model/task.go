package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	TaskId      primitive.ObjectID `bson:"task_id"`
	Status      string             `bson:"status"`
	LastUpdated time.Time          `bson:"last_updated"`
}

type TaskMaster struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	RewardAmount float32            `bson:"reward_amount"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
}

type TaskData struct {
	Task       Task       `bson:"tasks" json:"task"`
	TaskDetail TaskMaster `bson:"task_detail" json:"task_detail"`
}

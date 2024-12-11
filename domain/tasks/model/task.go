package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	TaskId      primitive.ObjectID `bson:"task_id" json:"task_id"`
	Status      string             `bson:"status" json:"status"`
	LastUpdated time.Time          `bson:"last_updated" json:"last_updated"`
}

type TaskMaster struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	RewardAmount float32            `bson:"reward_amount" json:"reward_amount"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	Type         string             `bson:"type" json:"type"`
}
type TaskData struct {
	Task       Task       `bson:"tasks" json:"task"`
	TaskDetail TaskMaster `bson:"task_detail" json:"task_detail"`
}

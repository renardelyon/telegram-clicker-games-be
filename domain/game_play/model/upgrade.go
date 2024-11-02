package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpgradeMaster struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	MaxLevel    int32              `bson:"max_level"`
	BaseCost    float64            `bson:"base_cost"`
	Description string             `bson:"description"`
	Effect      string             `bson:"effect"`
}

type Upgrade struct {
	UpgradeId  primitive.ObjectID `bson:"upgrade_id"`
	NextCost   float64            `bson:"next_cost"`
	Level      int32              `bson:"level"`
	AcquiredAt time.Time          `bson:"acquired_at"`
}

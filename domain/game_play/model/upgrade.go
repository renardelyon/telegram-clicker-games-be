package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpgradeMaster struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MaxLevel      int32              `bson:"max_level" json:"max_level"`
	BaseCost      float64            `bson:"base_cost" json:"base_cost"`
	Description   string             `bson:"description" json:"description"`
	Effect        string             `bson:"effect" json:"effect"`
	IncMultiplier float64            `bson:"inc_multiplier" json:"inc_multiplier"`
}

type Upgrade struct {
	UpgradeId  primitive.ObjectID `bson:"upgrade_id" json:"upgrade_id"`
	NextCost   float64            `bson:"next_cost" json:"next_cost"`
	Level      int32              `bson:"level" json:"level"`
	AcquiredAt time.Time          `bson:"acquired_at" json:"acquired_at"`
}

type UpgradeData struct {
	Upgrade       Upgrade       `bson:"upgrades" json:"upgrade"`
	UpgradeDetail UpgradeMaster `bson:"upgrade_detail" json:"upgrade_detail"`
}

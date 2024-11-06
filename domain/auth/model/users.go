package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	TelegramId int64              `bson:"telegram_id"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	UserName   string             `bson:"user_name"`
	UpdatedAt  *time.Time         `bson:"updated_at"`
	DeletedAt  *time.Time         `bson:"deleted_at"`
	CreatedAt  time.Time          `bson:"created_at"`
	LangCode   string             `bson:"language_code"`
	IsPremium  bool               `bson:"is_premium"`
	Referral   Referral           `bson:"referral"`
	GameStates GameState          `bson:"game_states"`
	Upgrades   []Upgrade          `bson:"upgrades"`
	Tasks      []Task             `bson:"tasks"`
}

type GameState struct {
	ClickCount       int64     `bson:"click_count"`
	Balance          float64   `bson:"balance"`
	TotalBalance     float64   `bson:"total_balance"`
	Energy           int32     `bson:"energy"`
	MaxEnergy        int32     `bson:"max_energy"`
	LastEnergyUpdate time.Time `bson:"last_energy_update"`
	MiningRate       float64   `bson:"mining_rate"`
}

type Referral struct {
	ReferredBy *int  `bson:"referred_by"` // telegram_id
	Referrals  []int `bson:"referrals"`
}

type Upgrade struct {
	UpgradeId  primitive.ObjectID `bson:"upgrade_id"`
	NextCost   float64            `bson:"next_cost"`
	Level      int32              `bson:"level"`
	AcquiredAt time.Time          `bson:"acquired_at"`
}

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

type UpgradeMaster struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	MaxLevel    int32              `bson:"max_level"`
	BaseCost    float64            `bson:"base_cost"`
	Description string             `bson:"description"`
	Effect      string             `bson:"effect"`
}

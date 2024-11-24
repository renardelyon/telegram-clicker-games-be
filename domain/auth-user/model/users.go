package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	TelegramId int64              `bson:"telegram_id" json:"telegram_id"`
	FirstName  string             `bson:"first_name" json:"first_name"`
	LastName   string             `bson:"last_name" json:"last_name"`
	UserName   string             `bson:"user_name" json:"user_name"`
	UpdatedAt  *time.Time         `bson:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time         `bson:"deleted_at" json:"deleted_at"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	LangCode   string             `bson:"language_code" json:"language_code"`
	IsPremium  bool               `bson:"is_premium" json:"is_premium"`
	Referral   Referral           `bson:"referral" json:"referral"`
	GameStates GameState          `bson:"game_states" json:"game_states"`
	Upgrades   []Upgrade          `bson:"upgrades" json:"upgrades"`
	Tasks      []Task             `bson:"tasks" json:"tasks"`
}

type GameState struct {
	ClickCount       int64     `bson:"click_count" json:"click_count"`
	Balance          float64   `bson:"balance" json:"balance"`
	TotalBalance     float64   `bson:"total_balance" json:"total_balance"`
	Energy           int32     `bson:"energy" json:"energy"`
	BaseEnergy       int32     `bson:"base_energy" json:"base_energy"`
	LastEnergyUpdate time.Time `bson:"last_energy_update" json:"last_energy_update"`
	MiningRate       float64   `bson:"mining_rate" json:"mining_rate"`
}

type Referral struct {
	ReferredBy *int  `bson:"referred_by" json:"referred_by"` // telegram_id
	Referrals  []int `bson:"referrals" json:"referrals"`
}

type Upgrade struct {
	UpgradeId  primitive.ObjectID `bson:"upgrade_id" json:"upgrade_id"`
	NextCost   float64            `bson:"next_cost" json:"next_cost"`
	Level      int32              `bson:"level" json:"level"`
	AcquiredAt time.Time          `bson:"acquired_at" json:"acquired_at"`
}

type Task struct {
	TaskId      primitive.ObjectID `bson:"task_id" json:"task_id"`
	Status      string             `bson:"status" json:"status"`
	LastUpdated time.Time          `bson:"last_updated" json:"last_updated"`
}

type TaskMaster struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RewardAmount float32            `bson:"reward_amount" json:"reward_amount"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
}

type UpgradeMaster struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MaxLevel      int32              `bson:"max_level" json:"max_level"`
	BaseCost      float64            `bson:"base_cost" json:"base_cost"`
	Description   string             `bson:"description" json:"description"`
	Effect        string             `bson:"effect" json:"effect"`
	IncMultiplier float64            `bson:"inc_multiplier" json:"inc_multiplier"`
}

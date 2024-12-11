package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	TelegramId int64              `bson:"telegram_id" json:"telegram_id"`
	FirstName  string             `bson:"first_name" json:"first_name"`
	LastName   string             `bson:"last_name" json:"last_name"`
	UserName   string             `bson:"user_name" json:"user_name"`
	PhotUrl    string             `bson:"photo_url" json:"photo_url"`
	GameStates struct {
		TotalBalance float64 `bson:"total_balance" json:"total_balance"`
	} `bson:"game_states" json:"game_states"`
}

type Referral struct {
	ReferredBy int   `bson:"referred_by"` // telegram_id
	Referrals  []int `bson:"referrals"`
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	TelegramId int64              `bson:"telegram_id"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	UserName   string             `bson:"user_name"`
}

type Referral struct {
	ReferredBy int   `bson:"referred_by"` // telegram_id
	Referrals  []int `bson:"referrals"`
}

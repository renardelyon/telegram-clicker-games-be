package model

import "time"

type GameState struct {
	ClickCount       int64     `bson:"click_count"`
	Balance          float64   `bson:"balance"`
	TotalBalance     float64   `bson:"total_balance"`
	Energy           int32     `bson:"energy"`
	BaseEnergy       int32     `bson:"base_energy"`
	LastEnergyUpdate time.Time `bson:"last_energy_update"`
	MiningRate       float64   `bson:"mining_rate"`
}

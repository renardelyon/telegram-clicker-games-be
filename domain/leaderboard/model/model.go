package model

type User struct {
	FirstName  string     `bson:"first_name" json:"first_name"`
	LastName   string     `bson:"last_name" json:"last_name"`
	GameStates GameStates `bson:"game_states" json:"game_states"`
}

type SortedUsers struct {
	Users []User `json:"users"`
}

type GameStates struct {
	TotalBalance float64 `bson:"total_balance" json:"balance"`
}

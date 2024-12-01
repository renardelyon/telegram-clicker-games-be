package payload

import (
	"telegram-clicker-game-be/pkg/utils"
)

type SubmitTapsPayload struct {
	Taps int            `json:"taps"`
	Time utils.JsonTime `json:"time"`
}

type BuyUpgradePayload struct {
	UpgradeId string `json:"upgrade_id"`
}

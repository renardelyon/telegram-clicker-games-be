package payload

type SubmitTapsPayload struct {
	Taps int `json:"taps"`
}

type BuyUpgradePayload struct {
	UpgradeId string `json:"upgrade_id"`
}

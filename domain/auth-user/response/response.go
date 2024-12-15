package response

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type TelegramMembershipResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		User        User   `json:"user"`
		Status      string `json:"status"`
		IsAnonymous bool   `json:"is_anonymous"`
	} `json:"result"`
}

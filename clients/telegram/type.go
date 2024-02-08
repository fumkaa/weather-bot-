package telegram

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int              `json:"update_id"`
	Message  *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	UserId   int    `json:"id"`
	UserName string `json:"username"`
}

type Chat struct {
	ChatId int `json:"id"`
}

package options

type WebhookDiscord struct {
	Content string `json:"content"`
}

type WebhookSlack struct {
	Text string `json:"text"`
}

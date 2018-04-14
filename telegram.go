package telegram

// BotAPI is telegram bot api structure
type BotAPI struct {
	Token string
}

// New will return new telegram bot api
func New(token string) *BotAPI {
	api := new(BotAPI)
	api.Token = token
	return api
}

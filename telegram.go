package telegram

// BotOption type.
type BotOption func(*Bot)

type Bot struct {
	token   string
	baseURL string
}

// New return a telegram bot instance.
func New(token string, options ...BotOption) *Bot {
	b := Bot{
		token:   token,
		baseURL: "https://api.telegram.org",
	}

	for i := range options {
		options[i](&b)
	}

	return &b
}

func (b Bot) GetToken() string {
	return b.token
}

// SetBaseURL method.
func SetBaseURL(v string) BotOption {
	return func(b *Bot) {
		b.baseURL = v
	}
}

func (b Bot) GetBaseURL() string {
	return b.baseURL
}

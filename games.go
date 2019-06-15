package telegram

type SendGameRequest interface {
	ChatID(int) SendGameRequest
	GameShortName(string) SendGameRequest
	DisableNotification() SendGameRequest
	ReplyToMessageID(int) SendGameRequest
	ReplyMarkup(InlineKeyboardMarkup) SendGameRequest
	Do() (*Message, error)
}

func (b *bot) SendGame() SendGameRequest {
	panic("implement me")
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

type CallbackGame interface{}

type SetGameScoreRequest interface {
	UserID(int) SetGameScoreRequest
	Score(int) SetGameScoreRequest
	Force() SetGameScoreRequest
	DisableEditMessage(int) SetGameScoreRequest
	ChatID(int) SetGameScoreRequest
	MessageID(int) SetGameScoreRequest
	InlineMessageID(string) SetGameScoreRequest
	Do() (*Message, bool, error)
}

func (b *bot) SetGameScore() SetGameScoreRequest {
	panic("implement me")
}

type GetGameHighScoresRequest interface {
	UserID(int) SetGameScoreRequest
	ChatID(int) SetGameScoreRequest
	MessageID(int) SetGameScoreRequest
	InlineMessageID(string) SetGameScoreRequest
	Do() ([]GameHighScore, error)
}

func (b *bot) GetGameHighScores() GetGameHighScoresRequest {
	panic("implement me")
}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

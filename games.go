package telegram

type SendGameResponse interface {
}

func (b *bot) SendGame(options ...Option) (SendGameResponse, error) {
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

type SetGameScoreResponse struct {
	Response
	Result *Message `json:"result,omitempty"`
}

func (b *bot) SetGameScore(...Option) (SetGameScoreResponse, error) {
	panic("implement me")
}

type GetGameHighScoresResponse struct {
	Response
	Result []GameHighScore `json:"result,omitempty"`
}

func (b *bot) GetGameHighScores(...Option) (GetGameHighScoresResponse, error) {
	panic("implement me")
}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

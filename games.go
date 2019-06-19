package telegram

type SendGameResponse interface {
	Response
	GetMessage() *Message
}

type sendGameResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendGameResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendGame(options ...Option) (SendGameResponse, error) {
	var res sendGameResponse
	err := doRequest(b.Token, "sendGame", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
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

type SetGameScoreResponse interface {
	Response
	GetEditedMessage() *Message
}

type setGameScoreResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *setGameScoreResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) SetGameScore(options ...Option) (SetGameScoreResponse, error) {
	var res setGameScoreResponse
	err := doRequest(b.Token, "setGameScore", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetGameHighScoresResponse interface {
	Response
	GetGameHighScores() []GameHighScore
}

type getGameHighScoresResponse struct {
	response
	Result []GameHighScore `json:"result,omitempty"`
}

func (r *getGameHighScoresResponse) GetGameHighScores() []GameHighScore {
	return r.Result
}

func (b *bot) GetGameHighScores(options ...Option) (GetGameHighScoresResponse, error) {
	var res getGameHighScoresResponse
	err := doRequest(b.Token, "getGameHighScores", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

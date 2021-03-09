package telegram

import "context"

// SendGameResponse interface
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

func (b *bot) SendGame(ctx context.Context, options ...MethodOption) (SendGameResponse, error) {
	var res sendGameResponse

	err := b.doRequest(ctx, "sendGame", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Game struct
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// CallbackGame interface
type CallbackGame interface{}

// SetGameScoreResponse interface
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

func (b *bot) SetGameScore(ctx context.Context, options ...MethodOption) (SetGameScoreResponse, error) {
	var res setGameScoreResponse

	err := b.doRequest(ctx, "setGameScore", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetGameHighScoresResponse interface
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

func (b *bot) GetGameHighScores(ctx context.Context, options ...MethodOption) (GetGameHighScoresResponse, error) {
	var res getGameHighScoresResponse

	err := b.doRequest(ctx, "getGameHighScores", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GameHighScore struct
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

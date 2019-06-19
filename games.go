package telegram

type SendGameOption func(*request)

type SendGameResponse struct {
	Response
	Result *Message `json:"result,omitempty"`
}

func (b *bot) SendGame(options ...SendGameOption) (*SendGameResponse, error) {
	req := newRequest(b.Token, "sendGame")

	for i := range options {
		options[i](req)
	}

	var res SendGameResponse
	err := req.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SendGameChatID(v int) SendGameOption {
	return func(r *request) {
		r.setInt("chat_id", v)
	}
}

func SendGameGameShortName(v string) SendGameOption {
	return func(r *request) {
		r.setString("game_short_name", v)
	}
}

func SendGameDisableNotification() SendGameOption {
	return func(r *request) {
		r.setString("disable_notification", "true")
	}
}

func SendGameReplyToMessageID(v int) SendGameOption {
	return func(r *request) {
		r.setInt("reply_to_message_id", v)
	}
}

func SendGameReplyMarkup(v InlineKeyboardMarkup) SendGameOption {
	return func(r *request) {
		r.setObject("reply_markup", v)
	}
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

type SetGameScoreOption func(*request)

type SetGameScoreResponse struct {
	Response
	Result *Message `json:"result,omitempty"`
}

func (b *bot) SetGameScore(...SetGameScoreOption) (*SetGameScoreResponse, error) {
	panic("implement me")
}

type GetGameHighScoresOption func(*request)

type GetGameHighScoresResponse struct {
	Response
	Result []GameHighScore `json:"result,omitempty"`
}

func (b *bot) GetGameHighScores(...GetGameHighScoresOption) (*GetGameHighScoresResponse, error) {
	panic("implement me")
}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

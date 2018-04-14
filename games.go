package telegram

// SendGame ...
func (obj *BotAPI) SendGame(chatID int32, gameShortName string, disableNotification *bool, replyToMessageID *int32, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: chatID,
		},
		{
			name:     "game_short_name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: gameShortName,
		},
		{
			name:     "disable_notification",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: disableNotification,
		},
		{
			name:     "reply_to_message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: replyToMessageID,
		},
		{
			name:     "reply_markup",
			required: false,
			types: []parameterType{
				parameterTypeInlineKeyboardMarkup,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendGame", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

// Game  ...
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []PhotoSize      `json:"photo"`
	Text         *string          `json:"text,omitempty"`
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

// Animation  ...
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName *string    `json:"file_name"`
	MimeType *string    `json:"mime_type"`
	FileSize *int32     `json:"file_size"`
}

// CallbackGame  ...
type CallbackGame interface{}

// SetGameScore ...
func (obj *BotAPI) SetGameScore(userID int32, score int32, force *bool, disableEditMessage *bool, chatID *int32, messageID *int32, inlineMessageID *string) (*Message, error) {

	parameters := []parameter{
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
		{
			name:     "score",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: score,
		},
		{
			name:     "force",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: force,
		},
		{
			name:     "disable_edit_message",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: disableEditMessage,
		},
		{
			name:     "chat_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: chatID,
		},
		{
			name:     "message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
		{
			name:     "inline_message_id",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: inlineMessageID,
		},
	}

	res, err := obj.callMethod("setGameScore", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

// GetGameHighScores ...
func (obj *BotAPI) GetGameHighScores(userID int32, chatID *int32, messageID *int32, inlineMessageID *string) (*[]GameHighScore, error) {

	parameters := []parameter{
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
		{
			name:     "chat_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: chatID,
		},
		{
			name:     "message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
		{
			name:     "inline_message_id",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: inlineMessageID,
		},
	}

	res, err := obj.callMethod("getGameHighScores", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*[]GameHighScore), nil
}

// GameHighScore  ...
type GameHighScore struct {
	Position int32 `json:"position"`
	User     User  `json:"user"`
	Score    int32 `json:"score"`
}

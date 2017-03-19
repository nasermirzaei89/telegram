package telegram

import "errors"

func (obj *API) sendGame(chatID int64, gameShortName string, disableNotification bool, replyToMessageID int64, replyMarkup *InlineKeyboardMarkup) (*Message, error) {
	return nil, errors.New("Not implemented")
}

// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

// Animation ...
// You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
type Animation struct {
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int64     `json:"file_size"`
}

// CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame interface{}

func (obj *API) setGameScore(userID int64, score int64, force bool, disableEditMessage bool, chatID int64, messageID int64, inlineMessageID string) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

func (obj *API) getGameHighScores(userID int64, chatID int64, messageID int64, inlineMessageID string) ([]GameHighScore, error) {
	return nil, errors.New("Not implemented")
}

// GameHighScore represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int64 `json:"position"`
	User     *User `json:"user"`
	Score    int64 `json:"score"`
}

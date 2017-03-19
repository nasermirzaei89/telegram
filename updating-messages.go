package telegram

import "errors"

func (obj *API) editMessageText(chatID interface{}, messageID int64, inlineMessageID int64, text string, parseMode bool, disableWebPagePreview bool, replyMarkup interface{}) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

func (obj *API) editMessageCaption(chatID interface{}, messageID int64, inlineMessageID int64, caption string, replyMarkup interface{}) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

func (obj *API) editMessageReplyMarkup(chatID interface{}, messageID int64, inlineMessageID int64, replyMarkup interface{}) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

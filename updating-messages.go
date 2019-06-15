package telegram

type EditMessageTextRequest interface {
	ChatID(int) EditMessageTextRequest
	ChatUsername(string) EditMessageTextRequest
	MessageID(int) EditMessageTextRequest
	InlineMessageID(int) EditMessageTextRequest
	Text(string) EditMessageTextRequest
	ParseMode(string) EditMessageTextRequest
	DisableWebPagePreview() EditMessageTextRequest
	ReplyMarkup(InlineKeyboardMarkup) EditMessageTextRequest
	Do() (*Message, bool, error)
}

func (b *bot) EditMessageText() EditMessageTextRequest {
	panic("implement me")
}

type EditMessageCaptionRequest interface {
	ChatID(int) EditMessageCaptionRequest
	ChatUsername(string) EditMessageCaptionRequest
	MessageID(int) EditMessageCaptionRequest
	InlineMessageID(int) EditMessageCaptionRequest
	Caption(string) EditMessageCaptionRequest
	ParseMode(string) EditMessageCaptionRequest
	ReplyMarkup(InlineKeyboardMarkup) EditMessageCaptionRequest
	Do() (*Message, bool, error)
}

func (b *bot) EditMessageCaption() EditMessageCaptionRequest {
	panic("implement me")
}

type EditMessageMediaRequest interface {
	ChatID(int) EditMessageMediaRequest
	ChatUsername(string) EditMessageMediaRequest
	MessageID(int) EditMessageMediaRequest
	InlineMessageID(int) EditMessageMediaRequest
	Media(InputMedia) EditMessageMediaRequest
	ReplyMarkup(InlineKeyboardMarkup) EditMessageMediaRequest
	Do() (*Message, bool, error)
}

func (b *bot) EditMessageMedia() EditMessageMediaRequest {
	panic("implement me")
}

type EditMessageReplyMarkupRequest interface {
	ChatID(int) EditMessageReplyMarkupRequest
	ChatUsername(string) EditMessageReplyMarkupRequest
	MessageID(int) EditMessageReplyMarkupRequest
	InlineMessageID(int) EditMessageReplyMarkupRequest
	ReplyMarkup(InlineKeyboardMarkup) EditMessageReplyMarkupRequest
	Do() (*Message, bool, error)
}

func (b *bot) EditMessageReplyMarkup() EditMessageReplyMarkupRequest {
	panic("implement me")
}

type StopPollRequest interface {
	ChatID(int) StopPollRequest
	ChatUsername(string) StopPollRequest
	MessageID(int) StopPollRequest
	ReplyMarkup(InlineKeyboardMarkup) StopPollRequest
	Do() (*Message, bool, error)
}

func (b *bot) StopPoll() StopPollRequest {
	panic("implement me")
}

type DeleteMessageRequest interface {
	ChatID(int) DeleteMessageRequest
	ChatUsername(string) DeleteMessageRequest
	MessageID(int) DeleteMessageRequest
	Do() (*Message, bool, error)
}

func (b *bot) DeleteMessage() DeleteMessageRequest {
	panic("implement me")
}

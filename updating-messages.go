package telegram

type EditMessageTextResponse interface {
}

func (b *bot) EditMessageText(options ...Option) (EditMessageTextResponse, error) {
	panic("implement me")
}

type EditMessageCaptionResponse interface {
}

func (b *bot) EditMessageCaption(options ...Option) (EditMessageCaptionResponse, error) {
	panic("implement me")
}

type EditMessageMediaResponse interface {
}

func (b *bot) EditMessageMedia(options ...Option) (EditMessageMediaResponse, error) {
	panic("implement me")
}

type EditMessageReplyMarkupResponse interface {
}

func (b *bot) EditMessageReplyMarkup(options ...Option) (EditMessageReplyMarkupResponse, error) {
	panic("implement me")
}

type StopPollResponse interface {
}

func (b *bot) StopPoll(options ...Option) (StopPollResponse, error) {
	panic("implement me")
}

type DeleteMessageResponse interface {
}

func (b *bot) DeleteMessage(options ...Option) (DeleteMessageResponse, error) {
	panic("implement me")
}

package telegram

type EditMessageTextResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageTextResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageTextResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageText(options ...Option) (EditMessageTextResponse, error) {
	var res editMessageTextResponse
	err := doRequest(b.Token, "editMessageText", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type EditMessageCaptionResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageCaptionResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageCaptionResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageCaption(options ...Option) (EditMessageCaptionResponse, error) {
	var res editMessageCaptionResponse
	err := doRequest(b.Token, "editMessageCaption", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type EditMessageMediaResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageMediaResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageMediaResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageMedia(options ...Option) (EditMessageMediaResponse, error) {
	var res editMessageMediaResponse
	err := doRequest(b.Token, "editMessageMedia", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type EditMessageReplyMarkupResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageReplyMarkupResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageReplyMarkupResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageReplyMarkup(options ...Option) (EditMessageReplyMarkupResponse, error) {
	var res editMessageReplyMarkupResponse
	err := doRequest(b.Token, "editMessageReplyMarkup", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type StopPollResponse interface {
	Response
	GetStoppedPoll() *Poll
}

type stopPollResponse struct {
	response
	Result *Poll `json:"result,omitempty"`
}

func (r *stopPollResponse) GetStoppedPoll() *Poll {
	return r.Result
}

func (b *bot) StopPoll(options ...Option) (StopPollResponse, error) {
	var res stopPollResponse
	err := doRequest(b.Token, "stopPoll", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type DeleteMessageResponse interface {
	Response
}

type deleteMessageResponse struct {
	response
}

func (b *bot) DeleteMessage(options ...Option) (DeleteMessageResponse, error) {
	var res deleteMessageResponse
	err := doRequest(b.Token, "deleteMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

package telegram

// EditMessageTextResponse interface
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

func (b *bot) EditMessageText(options ...MethodOption) (EditMessageTextResponse, error) {
	var res editMessageTextResponse
	err := b.doRequest("editMessageText", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageCaptionResponse interface
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

func (b *bot) EditMessageCaption(options ...MethodOption) (EditMessageCaptionResponse, error) {
	var res editMessageCaptionResponse
	err := b.doRequest("editMessageCaption", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageMediaResponse interface
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

func (b *bot) EditMessageMedia(options ...MethodOption) (EditMessageMediaResponse, error) {
	var res editMessageMediaResponse
	err := b.doRequest("editMessageMedia", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageReplyMarkupResponse interface
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

func (b *bot) EditMessageReplyMarkup(options ...MethodOption) (EditMessageReplyMarkupResponse, error) {
	var res editMessageReplyMarkupResponse
	err := b.doRequest("editMessageReplyMarkup", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// StopPollResponse interface
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

func (b *bot) StopPoll(options ...MethodOption) (StopPollResponse, error) {
	var res stopPollResponse
	err := b.doRequest("stopPoll", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteMessageResponse interface
type DeleteMessageResponse interface {
	Response
}

type deleteMessageResponse struct {
	response
}

func (b *bot) DeleteMessage(options ...MethodOption) (DeleteMessageResponse, error) {
	var res deleteMessageResponse
	err := b.doRequest("deleteMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

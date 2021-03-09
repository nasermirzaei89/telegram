package telegram

import "context"

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

func (b *bot) EditMessageText(ctx context.Context, options ...MethodOption) (EditMessageTextResponse, error) {
	var res editMessageTextResponse

	err := b.doRequest(ctx, "editMessageText", &res, options...)
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

func (b *bot) EditMessageCaption(ctx context.Context, options ...MethodOption) (EditMessageCaptionResponse, error) {
	var res editMessageCaptionResponse

	err := b.doRequest(ctx, "editMessageCaption", &res, options...)
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

func (b *bot) EditMessageMedia(ctx context.Context, options ...MethodOption) (EditMessageMediaResponse, error) {
	var res editMessageMediaResponse

	err := b.doRequest(ctx, "editMessageMedia", &res, options...)
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

func (b *bot) EditMessageReplyMarkup(ctx context.Context, options ...MethodOption) (EditMessageReplyMarkupResponse, error) {
	var res editMessageReplyMarkupResponse

	err := b.doRequest(ctx, "editMessageReplyMarkup", &res, options...)
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

func (b *bot) StopPoll(ctx context.Context, options ...MethodOption) (StopPollResponse, error) {
	var res stopPollResponse

	err := b.doRequest(ctx, "stopPoll", &res, options...)
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

func (b *bot) DeleteMessage(ctx context.Context, options ...MethodOption) (DeleteMessageResponse, error) {
	var res deleteMessageResponse

	err := b.doRequest(ctx, "deleteMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

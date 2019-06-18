package telegram

import (
	"fmt"
	"strings"
)

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
}

type GetUpdatesRequest interface {
	Offset(int) GetUpdatesRequest
	Limit(int) GetUpdatesRequest
	Timeout(int) GetUpdatesRequest
	AllowedUpdates(...string) GetUpdatesRequest
	Do() (*GetUpdatesResponse, error)
}

type GetUpdatesResponse struct {
	Response
	Result []Update `json:"result,omitempty"`
}

func (b *bot) GetUpdates() GetUpdatesRequest {
	return &getUpdatesRequest{*newRequest(b.Token, "getUpdates")}
}

type getUpdatesRequest struct {
	request
}

func (r *getUpdatesRequest) Offset(v int) GetUpdatesRequest {
	err := r.writer.WriteField("offset", fmt.Sprintf("%d", v))
	if err != nil {
		r.err = err
	}

	return r
}

func (r *getUpdatesRequest) Limit(v int) GetUpdatesRequest {
	err := r.writer.WriteField("limit", fmt.Sprintf("%d", v))
	if err != nil {
		r.err = err
	}

	return r
}

func (r *getUpdatesRequest) Timeout(v int) GetUpdatesRequest {
	err := r.writer.WriteField("timeout", fmt.Sprintf("%d", v))
	if err != nil {
		r.err = err
	}

	return r
}

func (r *getUpdatesRequest) AllowedUpdates(v ...string) GetUpdatesRequest {
	str := "[]"
	if len(v) == 0 {
		str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
	}
	err := r.writer.WriteField("allowed_updates", str)
	if err != nil {
		r.err = err
	}

	return r
}

func (r *getUpdatesRequest) Do() (*GetUpdatesResponse, error) {
	var res GetUpdatesResponse
	err := r.request.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetWebhookRequest interface {
	URL(string) SetWebhookRequest
	Certificate(InputFile) SetWebhookRequest
	MaxConnections(string) SetWebhookRequest
	AllowedUpdates(...string) SetWebhookRequest
	Do() (*SetWebhookResponse, error)
}

type SetWebhookResponse struct {
	Response
	Result bool `json:"result,omitempty"`
}

func (b *bot) SetWebhook() SetWebhookRequest {
	return &setWebhookRequest{*newRequest(b.Token, "setWebhook")}
}

type setWebhookRequest struct {
	request
}

func (r *setWebhookRequest) URL(v string) SetWebhookRequest {
	err := r.writer.WriteField("url", v)
	if err != nil {
		r.err = err
	}

	return r
}

func (r *setWebhookRequest) Certificate(v InputFile) SetWebhookRequest {
	panic("implement me")
}

func (r *setWebhookRequest) MaxConnections(v string) SetWebhookRequest {
	err := r.writer.WriteField("max_connections", v)
	if err != nil {
		r.err = err
	}

	return r
}

func (r *setWebhookRequest) AllowedUpdates(v ...string) SetWebhookRequest {
	str := "[]"
	if len(v) == 0 {
		str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
	}
	err := r.writer.WriteField("allowed_updates", str)
	if err != nil {
		r.err = err
	}

	return r
}

func (r *setWebhookRequest) Do() (*SetWebhookResponse, error) {
	var res SetWebhookResponse
	err := r.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type DeleteWebhookRequest interface {
	Do() (*DeleteWebhookResponse, error)
}

type DeleteWebhookResponse struct {
	Response
	Result bool `json:"result,omitempty"`
}

func (b *bot) DeleteWebhook() DeleteWebhookRequest {
	return &deleteWebhookRequest{*newRequest(b.Token, "deleteWebhook")}
}

type deleteWebhookRequest struct {
	request
}

func (r *deleteWebhookRequest) Do() (*DeleteWebhookResponse, error) {
	var res DeleteWebhookResponse
	err := r.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetWebhookInfoRequest interface {
	Do() (*GetWebhookInfoResponse, error)
}

type GetWebhookInfoResponse struct {
	Response
	Result *WebhookInfo `json:"result,omitempty"`
}

func (b *bot) GetWebhookInfo() GetWebhookInfoRequest {
	return &getWebhookInfoRequest{*newRequest(b.Token, "getWebhookInfo")}
}

type getWebhookInfoRequest struct {
	request
}

func (r *getWebhookInfoRequest) Do() (*GetWebhookInfoResponse, error) {
	var res GetWebhookInfoResponse
	err := r.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        *int     `json:"last_error_date,omitempty"`
	LastErrorMessage     *string  `json:"last_error_message,omitempty"`
	MaxConnections       *int     `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

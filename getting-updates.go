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

type GetUpdatesOption func(*request)

type GetUpdatesResponse struct {
	Response
	Result []Update `json:"result,omitempty"`
}

func (b *bot) GetUpdates(options ...GetUpdatesOption) (*GetUpdatesResponse, error) {
	req := newRequest(b.Token, "getUpdates")

	for i := range options {
		options[i](req)
	}

	var res GetUpdatesResponse
	err := req.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetUpdatesOffset(v int) GetUpdatesOption {
	return func(r *request) {
		err := r.writer.WriteField("offset", fmt.Sprintf("%d", v))
		if err != nil {
			r.err = err
		}
	}
}

func GetUpdatesLimit(v int) GetUpdatesOption {
	return func(r *request) {
		err := r.writer.WriteField("limit", fmt.Sprintf("%d", v))
		if err != nil {
			r.err = err
		}
	}
}

func GetUpdatesTimeout(v int) GetUpdatesOption {
	return func(r *request) {
		err := r.writer.WriteField("timeout", fmt.Sprintf("%d", v))
		if err != nil {
			r.err = err
		}
	}
}

func GetUpdatesAllowedUpdates(v ...string) GetUpdatesOption {
	return func(r *request) {
		str := "[]"
		if len(v) == 0 {
			str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
		}
		err := r.writer.WriteField("allowed_updates", str)
		if err != nil {
			r.err = err
		}
	}
}

type SetWebhookOption func(*request)

type SetWebhookResponse struct {
	Response
	Result bool `json:"result,omitempty"`
}

func (b *bot) SetWebhook(options ...SetWebhookOption) (*SetWebhookResponse, error) {
	req := newRequest(b.Token, "setWebhook")

	for i := range options {
		options[i](req)
	}

	var res SetWebhookResponse
	err := req.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SetWebhookURL(v string) SetWebhookOption {
	return func(r *request) {
		err := r.writer.WriteField("url", v)
		if err != nil {
			r.err = err
		}
	}
}

func SetWebhookCertificate(v InputFile) SetWebhookOption {
	return func(r *request) {
		panic("implement me")
	}
}

func SetWebhookMaxConnections(v string) SetWebhookOption {
	return func(r *request) {
		err := r.writer.WriteField("max_connections", v)
		if err != nil {
			r.err = err
		}
	}
}

func SetWebhookAllowedUpdates(v ...string) SetWebhookOption {
	return func(r *request) {
		str := "[]"
		if len(v) == 0 {
			str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
		}
		err := r.writer.WriteField("allowed_updates", str)
		if err != nil {
			r.err = err
		}
	}
}

type DeleteWebhookResponse struct {
	Response
	Result bool `json:"result,omitempty"`
}

func (b *bot) DeleteWebhook() (*DeleteWebhookResponse, error) {
	req := newRequest(b.Token, "deleteWebhook")

	var res DeleteWebhookResponse
	err := req.do(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetWebhookInfoResponse struct {
	Response
	Result *WebhookInfo `json:"result,omitempty"`
}

func (b *bot) GetWebhookInfo() (*GetWebhookInfoResponse, error) {
	req := newRequest(b.Token, "getWebhookInfo")

	var res GetWebhookInfoResponse
	err := req.do(&res)
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

package telegram

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

type GetUpdatesResponse interface {
	Response
	GetUpdates() []Update
}

type getUpdatesResponse struct {
	response
	Result []Update `json:"result,omitempty"`
}

func (r *getUpdatesResponse) GetUpdates() []Update {
	return r.Result
}

func (b *bot) GetUpdates(options ...Option) (GetUpdatesResponse, error) {
	var res getUpdatesResponse
	err := doRequest(b.Token, "getUpdates", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetWebhookResponse interface {
	Response
}

type setWebhookResponse struct {
	response
}

func (b *bot) SetWebhook(options ...Option) (SetWebhookResponse, error) {
	var res setWebhookResponse
	err := doRequest(b.Token, "setWebhook", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type DeleteWebhookResponse interface {
	Response
}

type deleteWebhookResponse struct {
	response
}

func (b *bot) DeleteWebhook() (DeleteWebhookResponse, error) {
	var res deleteWebhookResponse
	err := doRequest(b.Token, "deleteWebhook", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetWebhookInfoResponse interface {
	Response
	GetWebhookInfo() *WebhookInfo
}

type getWebhookInfoResponse struct {
	response
	Result *WebhookInfo `json:"result,omitempty"`
}

func (r *getWebhookInfoResponse) GetWebhookInfo() *WebhookInfo {
	return r.Result
}

func (b *bot) GetWebhookInfo() (GetWebhookInfoResponse, error) {
	var res getWebhookInfoResponse
	err := doRequest(b.Token, "getWebhookInfo", &res)
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

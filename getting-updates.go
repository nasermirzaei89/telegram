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

type GetUpdatesRequest interface {
	Offset(int) GetUpdatesRequest
	Limit(int) GetUpdatesRequest
	Timeout(int) GetUpdatesRequest
	AllowedUpdates(...string) GetUpdatesRequest
	Do() ([]Update, error)
}

func (b *bot) GetUpdates() GetUpdatesRequest {
	panic("implement me")
}

type SetWebhookRequest interface {
	URL(string) SetWebhookRequest
	Certificate(InputFile) SetWebhookRequest
	MaxConnections(string) SetWebhookRequest
	AllowedUpdates(...string) SetWebhookRequest
	Do() (bool, error)
}

func (b *bot) SetWebhook() SetWebhookRequest {
	panic("implement me")
}

type DeleteWebhookRequest interface {
	Do() (bool, error)
}

func (b *bot) DeleteWebhook() DeleteWebhookRequest {
	panic("implement me")
}

type GetWebhookInfoRequest interface {
	Do() (*WebhookInfo, error)
}

func (b *bot) GetWebhookInfo() GetWebhookInfoRequest {
	panic("implement me")
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

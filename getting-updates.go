package telegram

type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
}

func (obj *BotAPI) GetUpdates(offset, limit, timeout *int64, allowedUpdates *[]string) ([]Update, error) {
	parameters := []parameter{
		{
			name:     "offset",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: offset,
		},
		{
			name:     "limit",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: limit,
		},
		{
			name:     "timeout",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: timeout,
		},
		{
			name:     "allowed_updates",
			required: false,
			types: []parameterType{
				parameterTypeArrayOfString,
			},
			value: allowedUpdates,
		},
	}

	res, err := obj.callMethod("getUpdates", parameters...)
	if err != nil {
		return nil, err
	}

	return res.([]Update), nil
}

func (obj *BotAPI) SetWebhook(url string, certificate *InputFile, maxConnections *int32, allowedUpdates *[]string) (*bool, error) {
	parameters := []parameter{
		{
			name:     "url",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: url,
		},
		{
			name:     "certificate",
			required: false,
			types: []parameterType{
				parameterTypeInputFile,
			},
			value: certificate,
		},
		{
			name:     "max_connections",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: maxConnections,
		},
		{
			name:     "allowed_updates",
			required: false,
			types: []parameterType{
				parameterTypeArrayOfString,
			},
			value: allowedUpdates,
		},
	}

	res, err := obj.callMethod("setWebhook", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) DeleteWebhook() (*bool, error) {
	res, err := obj.callMethod("deleteWebhook")
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) GetWebhookInfo() (*WebhookInfo, error) {
	res, err := obj.callMethod("getWebhookInfo")
	if err != nil {
		return nil, err
	}

	return res.(*WebhookInfo), nil
}

type WebhookInfo struct {
	URL                  string    `json:"url"`
	HasCustomCertificate bool      `json:"has_custom_certificate"`
	PendingUpdateCount   int32     `json:"pending_update_count"`
	LastErrorDate        *int32    `json:"last_error_date,omitempty"`
	LastErrorMessage     *string   `json:"last_error_message,omitempty"`
	MaxConnections       *int32    `json:"max_connections,omitempty"`
	AllowedUpdates       *[]string `json:"allowed_updates,omitempty"`
}

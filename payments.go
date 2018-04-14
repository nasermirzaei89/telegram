package telegram

func (obj *BotAPI) SendInvoice(chatID int32, title string, description string, payload string, providerToken string, startParameter string, currency string, prices []LabeledPrice, providerData *string, photoURL *string, photoSize *int32, photoWidth *int32, photoHeight *int32, needName *bool, needPhoneNumber *bool, needEmail *bool, needShippingAddress *bool, sendPhoneNumberToProvider *bool, sendEmailToProvider *bool, isFlexible *bool, disableNotification *bool, replyToMessageID *int32, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: chatID,
		},
		{
			name:     "title",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: title,
		},
		{
			name:     "description",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: description,
		},
		{
			name:     "payload",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: payload,
		},
		{
			name:     "provider_token",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: providerToken,
		},
		{
			name:     "start_parameter",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: startParameter,
		},
		{
			name:     "currency",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: currency,
		},
		{
			name:     "prices",
			required: true,
			types: []parameterType{
				parameterTypeArrayOfLabeledPrice,
			},
			value: prices,
		},
		{
			name:     "provider_data",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: providerData,
		},
		{
			name:     "photo_url",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: photoURL,
		},
		{
			name:     "photo_size",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: photoSize,
		},
		{
			name:     "photo_width",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: photoWidth,
		},
		{
			name:     "photo_height",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: photoHeight,
		},
		{
			name:     "need_name",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: needName,
		},
		{
			name:     "need_phone_number",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: needPhoneNumber,
		},
		{
			name:     "need_email",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: needEmail,
		},
		{
			name:     "need_shipping_address",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: needShippingAddress,
		},
		{
			name:     "send_phone_number_to_provider",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: sendPhoneNumberToProvider,
		},
		{
			name:     "send_email_to_provider",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: sendEmailToProvider,
		},
		{
			name:     "is_flexible",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: isFlexible,
		},
		{
			name:     "disable_notification",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: disableNotification,
		},
		{
			name:     "reply_to_message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: replyToMessageID,
		},
		{
			name:     "reply_markup",
			required: false,
			types: []parameterType{
				parameterTypeInlineKeyboardMarkup,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendInvoice", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) AnswerShippingQuery(shippingQueryID string, ok bool, shippingOptions *[]ShippingOption, errorMessage *string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "shipping_query_id",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: shippingQueryID,
		},
		{
			name:     "ok",
			required: true,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: ok,
		},
		{
			name:     "shipping_options",
			required: false,
			types: []parameterType{
				parameterTypeArrayOfShippingOption,
			},
			value: shippingOptions,
		},
		{
			name:     "error_message",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: errorMessage,
		},
	}

	res, err := obj.callMethod("answerShippingQuery", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, errorMessage *string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "pre_checkout_query_id",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: preCheckoutQueryID,
		},
		{
			name:     "ok",
			required: true,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: ok,
		},
		{
			name:     "error_message",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: errorMessage,
		},
	}

	res, err := obj.callMethod("answerPreCheckoutQuery", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int32  `json:"amount"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int32  `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	Email           *string          `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int32      `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        *string    `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string
}

type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int32      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

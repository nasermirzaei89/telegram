package telegram

type SendInvoiceResponse interface {
	Response
	GetMessage() *Message
}

type sendInvoiceResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendInvoiceResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendInvoice(options ...Option) (SendInvoiceResponse, error) {
	var res sendInvoiceResponse
	err := doRequest(b.Token, "sendInvoice", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type AnswerShippingQueryResponse interface {
	Response
}

type answerShippingQueryResponse struct {
	response
}

func (b *bot) AnswerShippingQuery(options ...Option) (AnswerShippingQueryResponse, error) {
	var res answerShippingQueryResponse
	err := doRequest(b.Token, "answerShippingQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type AnswerPreCheckoutQueryResponse interface {
	Response
}

type answerPreCheckoutQueryResponse struct {
	response
}

func (b *bot) AnswerPreCheckoutQuery(options ...Option) (AnswerPreCheckoutQueryResponse, error) {
	var res answerPreCheckoutQueryResponse
	err := doRequest(b.Token, "answerPreCheckoutQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
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
	TotalAmount             int        `json:"total_amount"`
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
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

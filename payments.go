package telegram

type SendInvoiceResponse interface {
}

func (b *bot) SendInvoice(options ...Option) (SendInvoiceResponse, error) {
	panic("implement me")
}

type AnswerShippingQueryResponse interface {
}

func (b *bot) AnswerShippingQuery(options ...Option) (AnswerShippingQueryResponse, error) {
	panic("implement me")
}

type AnswerPreCheckoutQueryResponse interface {
}

func (b *bot) AnswerPreCheckoutQuery(options ...Option) (AnswerPreCheckoutQueryResponse, error) {
	panic("implement me")
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

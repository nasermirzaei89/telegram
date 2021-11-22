package telegram

import "context"

type ChatType string

const (
	ChatTypeSender     ChatType = "sender"
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// InlineQuery struct
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *ChatType `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// AnswerInlineQueryResponse interface
type AnswerInlineQueryResponse interface {
	Response
}

type answerInlineQueryResponse struct {
	response
}

func (b *bot) AnswerInlineQuery(ctx context.Context, options ...MethodOption) (AnswerInlineQueryResponse, error) {
	var res answerInlineQueryResponse

	err := b.doRequest(ctx, "answerInlineQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// InlineQueryResult interface
type InlineQueryResult interface{}

// InlineQueryResultArticle struct
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	HideURL             *bool                 `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto struct
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          *int                  `json:"photo_width,omitempty"`
	PhotoHeight         *int                  `json:"photo_height,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif struct
type InlineQueryResultGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifURL              string                `json:"gif_url"`
	GifWidth            *int                  `json:"gif_width,omitempty"`
	GifHeight           *int                  `json:"gif_height,omitempty"`
	GifDuration         *int                  `json:"gif_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif struct
type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4URL            string                `json:"mpeg4_url"`
	Mpeg4Width          *int                  `json:"mpeg4_width,omitempty"`
	Mpeg4Height         *int                  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       *int                  `json:"mpeg4_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo struct
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	VideoWidth          *int                  `json:"video_width,omitempty"`
	VideoHeight         *int                  `json:"video_height,omitempty"`
	VideoDuration       *int                  `json:"video_duration,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio struct
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice struct
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	VoiceDuration       *string               `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument struct
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

// InlineQueryResultLocation struct
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   *float32              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int                  `json:"live_period,omitempty"`
	Heading              *int                  `json:"heading,omitempty"`
	ProximityAlertRadius *int                  `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL             *string               `json:"thumb_url,omitempty"`
	ThumbWidth           *int                  `json:"thumb_width,omitempty"`
	ThumbHeight          *int                  `json:"thumb_height,omitempty"`
}

// InlineQueryResultVenue struct
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	FoursquareType      *string               `json:"foursquare_type,omitempty"`
	GooglePlaceID       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact struct
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	VCard               *string               `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

// InlineQueryResultGame struct
type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQueryResultCachedPhoto struct
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif struct
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifFileID           string                `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif struct
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4FileID         string                `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker struct
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument struct
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo struct
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice struct
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio struct
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InputMessageContent interface
type InputMessageContent interface{}

// InputTextMessageContent struct
type InputTextMessageContent struct {
	MessageText           string  `json:"message_text"`
	ParseMode             *string `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool   `json:"disable_web_page_preview,omitempty"`
}

// InputLocationMessageContent struct
type InputLocationMessageContent struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

// InputVenueMessageContent struct
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

// InputContactMessageContent struct
type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

// InputInvoiceMessageContent represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  Currency       `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              *int           `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoURL                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int           `json:"photo_size,omitempty"`
	PhotoWidth                *int           `json:"photo_width,omitempty"`
	PhotoHeight               *int           `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
}

// ChosenInlineResult struct
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

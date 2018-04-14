package telegram

// InlineQuery ...
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// AnswerInlineQuery ...
func (obj *BotAPI) AnswerInlineQuery(inlineQueryID string, results []InlineQueryResult, cacheTime *int32, isPersonal *bool, nextOffset, switchPmText, switchPmParameter *string) (*bool, error) {
	parameters := []parameter{
		{
			name:     "inline_query_id",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: inlineQueryID,
		},
		{
			name:     "results",
			required: true,
			types: []parameterType{
				parameterTypeArrayOfInlineQueryResult,
			},
			value: results,
		},
		{
			name:     "cache_time",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: cacheTime,
		},
		{
			name:     "is_personal",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: isPersonal,
		},
		{
			name:     "next_offset",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: nextOffset,
		},
		{
			name:     "switch_pm_text",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: switchPmText,
		},
		{
			name:     "switch_pm_parameter",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: switchPmParameter,
		},
	}

	res, err := obj.callMethod("answerInlineQuery", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

// InlineQueryResult represents one result of an inline query.
type InlineQueryResult interface{}

// InlineQueryResultArticle represents a link to an article or web page.
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
	ThumbWidth          *int32                `json:"thumb_width,omitempty"`
	ThumbHeight         *int32                `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          *int32                `json:"photo_width,omitempty"`
	PhotoHeight         *int32                `json:"photo_height,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifURL              string                `json:"gif_url"`
	GifWidth            *int32                `json:"gif_width,omitempty"`
	GifHeight           *int32                `json:"gif_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4URL            string                `json:"mpeg4_url"`
	Mpeg4Width          *int32                `json:"mpeg4_width,omitempty"`
	Mpeg4Height         *int32                `json:"mpeg4_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VideoWidth          *int32                `json:"video_width,omitempty"`
	VideoHeight         *int32                `json:"video_height,omitempty"`
	VideoDuration       *int32                `json:"video_duration,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int32                `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VoiceDuration       *string               `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int32                `json:"thumb_width,omitempty"`
	ThumbHeight         *int32                `json:"thumb_height,omitempty"`
}

// InlineQueryResultLocation represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int32                `json:"thumb_width,omitempty"`
	ThumbHeight         *int32                `json:"thumb_height,omitempty"`
}

// InlineQueryResultVenue represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int32                `json:"thumb_width,omitempty"`
	ThumbHeight         *int32                `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int32                `json:"thumb_width,omitempty"`
	ThumbHeight         *int32                `json:"thumb_height,omitempty"`
}

// InlineQueryResultGame represents a Game.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifFileID           string                `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4FileID         string                `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InputMessageContent represents the content of a message to be sent as a result of an inline query
type InputMessageContent interface{}

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string  `json:"message_text"`
	ParseMode             *string `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool   `json:"disable_web_page_preview,omitempty"`
}

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID *string `json:"foursquare_id,omitempty"`
}

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

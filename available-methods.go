package telegram

func (obj *BotAPI) GetMe() (*User, error) {

	res, err := obj.callMethod("getMe")
	if err != nil {
		return nil, err
	}

	return res.(*User), nil
}

func (obj *BotAPI) SendMessage(chatID interface{}, text string, parseMode *string, disableWebPagePreview *bool, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "text",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: text,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
		},
		{
			name:     "disable_web_page_preview",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: disableWebPagePreview,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendMessage", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) ForwardMessage(chatID interface{}, fromChatID interface{}, disableNotification *bool, messageID int32) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "from_chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: fromChatID,
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
			name:     "message_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
	}

	res, err := obj.callMethod("forwardMessage", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendPhoto(chatID interface{}, photo interface{}, caption *string, parseMode *string, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "photo",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: photo,
		},
		{
			name:     "caption",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: caption,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendPhoto", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendAudio(chatID interface{}, audio interface{}, caption *string, parseMode *string, duration *int32, performer *string, title *string, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "audio",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: audio,
		},
		{
			name:     "caption",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: caption,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
		},
		{
			name:     "duration",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: duration,
		},
		{
			name:     "performer",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: performer,
		},
		{
			name:     "title",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: title,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendAudio", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendDocument(chatID interface{}, document interface{}, caption *string, parseMode *string, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "document",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: document,
		},
		{
			name:     "caption",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: caption,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendDocument", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendVideo(chatID interface{}, video interface{}, duration *int32, width *int32, height *int32, caption *string, parseMode *string, supportsStreaming *bool, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "video",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: video,
		},
		{
			name:     "duration",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: duration,
		},
		{
			name:     "width",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: width,
		},
		{
			name:     "height",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: height,
		},
		{
			name:     "caption",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: caption,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
		},
		{
			name:     "supports_streaming",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: supportsStreaming,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendVideo", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendVoice(chatID interface{}, voice interface{}, caption *string, parseMode *string, duration *int32, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "voice",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: voice,
		},
		{
			name:     "caption",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: caption,
		},
		{
			name:     "parse_mode",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: parseMode,
		},
		{
			name:     "duration",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: duration,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendVoice", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendVideoNote(chatID interface{}, videoNote interface{}, duration *int32, length *int32, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "video_note",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: videoNote,
		},
		{
			name:     "duration",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: duration,
		},
		{
			name:     "length",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: length,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendVideoNote", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendMediaGroup(chatID interface{}, media []InputMedia, disableNotification *bool, replyToMessageID *int32) (*[]Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "media",
			required: true,
			types: []parameterType{
				parameterTypeArrayOfInputMedia,
			},
			value: media,
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
	}

	res, err := obj.callMethod("sendMediaGroup", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*[]Message), nil
}

func (obj *BotAPI) SendLocation(chatID interface{}, latitude float32, longitude float32, livePeriod *int32, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "latitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: latitude,
		},
		{
			name:     "longitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: longitude,
		},
		{
			name:     "live_period",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: livePeriod,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendLocation", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) EditMessageLiveLocation(chatID *interface{}, messageID *int32, inlineMessageID *string, latitude float32, longitude float32, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
		{
			name:     "inline_message_id",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: inlineMessageID,
		},
		{
			name:     "latitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: latitude,
		},
		{
			name:     "longitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: longitude,
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

	res, err := obj.callMethod("editMessageLiveLocation", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) StopMessageLiveLocation(chatID *interface{}, messageID *int32, inlineMessageID *string, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "message_id",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
		{
			name:     "inline_message_id",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: inlineMessageID,
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

	res, err := obj.callMethod("stopMessageLiveLocation", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendVenue(chatID interface{}, latitude float32, longitude float32, title string, address string, foursquareID *string, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "latitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: latitude,
		},
		{
			name:     "longitude",
			required: true,
			types: []parameterType{
				parameterTypeFloatNumber,
			},
			value: longitude,
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
			name:     "address",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: address,
		},
		{
			name:     "foursquare_id",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: foursquareID,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendVenue", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendContact(chatID interface{}, phoneNumber string, firstName string, lastName *string, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "phone_number",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: phoneNumber,
		},
		{
			name:     "first_name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: firstName,
		},
		{
			name:     "last_name",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: lastName,
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
				parameterTypeReplyKeyboardMarkup,
				parameterTypeReplyKeyboardRemove,
				parameterTypeForceReply,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("sendContact", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) SendChatAction(chatID interface{}, action string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "action",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: action,
		},
	}

	res, err := obj.callMethod("sendChatAction", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) GetUserProfilePhotos(userID int32, offset *int32, limit *int32) (*UserProfilePhotos, error) {

	parameters := []parameter{
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
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
	}

	res, err := obj.callMethod("getUserProfilePhotos", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*UserProfilePhotos), nil
}

func (obj *BotAPI) GetFile(fileID string) (*File, error) {

	parameters := []parameter{
		{
			name:     "file_id",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: fileID,
		},
	}

	res, err := obj.callMethod("getFile", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*File), nil
}

func (obj *BotAPI) KickChatMember(chatID interface{}, userID int32, untilDate *int32) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
		{
			name:     "until_date",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: untilDate,
		},
	}

	res, err := obj.callMethod("kickChatMember", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) UnbanChatMember(chatID interface{}, userID int32) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
	}

	res, err := obj.callMethod("unbanChatMember", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) RestrictChatMember(chatID interface{}, userID int32, untilDate *int32, canSendMessages *bool, canSendMediaMessages *bool, canSendOtherMessages *bool, canAddWebPagePreviews *bool) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
		{
			name:     "until_date",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: untilDate,
		},
		{
			name:     "can_send_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canSendMessages,
		},
		{
			name:     "can_send_media_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canSendMediaMessages,
		},
		{
			name:     "can_send_other_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canSendOtherMessages,
		},
		{
			name:     "can_add_web_page_previews",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canAddWebPagePreviews,
		},
	}

	res, err := obj.callMethod("restrictChatMember", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) PromoteChatMember(chatID interface{}, userID int32, canChangeInfo *bool, canPostMessages *bool, canEditMessages *bool, canDeleteMessages *bool, canInviteUsers *bool, canRestrictMembers *bool, canPinMessages *bool, canPromoteMembers *bool) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
		{
			name:     "can_change_info",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canChangeInfo,
		},
		{
			name:     "can_post_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canPostMessages,
		},
		{
			name:     "can_edit_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canEditMessages,
		},
		{
			name:     "can_delete_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canDeleteMessages,
		},
		{
			name:     "can_invite_users",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canInviteUsers,
		},
		{
			name:     "can_restrict_members",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canRestrictMembers,
		},
		{
			name:     "can_pin_messages",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canPinMessages,
		},
		{
			name:     "can_promote_members",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: canPromoteMembers,
		},
	}

	res, err := obj.callMethod("promoteChatMember", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) ExportChatInviteLink(chatID interface{}) (*string, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("exportChatInviteLink", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*string), nil
}

func (obj *BotAPI) SetChatPhoto(chatID interface{}, photo InputFile) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "photo",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
			},
			value: photo,
		},
	}

	res, err := obj.callMethod("setChatPhoto", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) DeleteChatPhoto(chatID interface{}) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("deleteChatPhoto", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) SetChatTitle(chatID interface{}, title string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
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
	}

	res, err := obj.callMethod("setChatTitle", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) SetChatDescription(chatID interface{}, description *string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "description",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: description,
		},
	}

	res, err := obj.callMethod("setChatDescription", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) PinChatMessage(chatID interface{}, messageID int32, disableNotification *bool) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "message_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: messageID,
		},
		{
			name:     "disable_notification",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: disableNotification,
		},
	}

	res, err := obj.callMethod("pinChatMessage", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) UnpinChatMessage(chatID interface{}) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("unpinChatMessage", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) LeaveChat(chatID interface{}) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("leaveChat", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) GetChat(chatID interface{}) (*Chat, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("getChat", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Chat), nil
}

func (obj *BotAPI) GetChatAdministrators(chatID interface{}) (*ChatMember, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("getChatAdministrators", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*ChatMember), nil
}

func (obj *BotAPI) GetChatMembersCount(chatID interface{}) (*int32, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("getChatMembersCount", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*int32), nil
}

func (obj *BotAPI) GetChatMember(chatID interface{}, userID int32) (*ChatMember, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "user_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: userID,
		},
	}

	res, err := obj.callMethod("getChatMember", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*ChatMember), nil
}

func (obj *BotAPI) SetChatStickerSet(chatID interface{}, stickerSetName string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
		{
			name:     "sticker_set_name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: stickerSetName,
		},
	}

	res, err := obj.callMethod("setChatStickerSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) DeleteChatStickerSet(chatID interface{}) (*bool, error) {

	parameters := []parameter{
		{
			name:     "chat_id",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
				parameterTypeString,
			},
			value: chatID,
		},
	}

	res, err := obj.callMethod("deleteChatStickerSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) AnswerCallbackQuery(callbackQueryID string, text *string, showAlert *bool, uRL *string, cacheTime *int32) (*bool, error) {

	parameters := []parameter{
		{
			name:     "callback_query_id",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: callbackQueryID,
		},
		{
			name:     "text",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: text,
		},
		{
			name:     "show_alert",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: showAlert,
		},
		{
			name:     "url",
			required: false,
			types: []parameterType{
				parameterTypeString,
			},
			value: uRL,
		},
		{
			name:     "cache_time",
			required: false,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: cacheTime,
		},
	}

	res, err := obj.callMethod("answerCallbackQuery", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

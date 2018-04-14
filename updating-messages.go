package telegram

// EditMessageText ...
func (obj *BotAPI) EditMessageText(chatID *interface{}, messageID *int32, inlineMessageID *string, text string, parseMode *string, disableWebPagePreview *bool, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

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
			name:     "reply_markup",
			required: false,
			types: []parameterType{
				parameterTypeInlineKeyboardMarkup,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("editMessageText", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

// EditMessageCaption ...
func (obj *BotAPI) EditMessageCaption(chatID *interface{}, messageID *int32, inlineMessageID *string, caption *string, parseMode *string, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

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
			name:     "reply_markup",
			required: false,
			types: []parameterType{
				parameterTypeInlineKeyboardMarkup,
			},
			value: replyMarkup,
		},
	}

	res, err := obj.callMethod("editMessageCaption", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

// EditMessageReplyMarkup ...
func (obj *BotAPI) EditMessageReplyMarkup(chatID *interface{}, messageID *int32, inlineMessageID *string, replyMarkup *InlineKeyboardMarkup) (*Message, error) {

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

	res, err := obj.callMethod("editMessageReplyMarkup", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

// DeleteMessage ...
func (obj *BotAPI) DeleteMessage(chatID interface{}, messageID int32) (*bool, error) {

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
	}

	res, err := obj.callMethod("deleteMessage", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

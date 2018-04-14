package telegram

type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int32         `json:"width"`
	Height       int32         `json:"height"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        *string       `json:"emoji,omitempty"`
	SetName      *string       `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     *int32        `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

func (obj *BotAPI) SendSticker(chatID interface{}, sticker interface{}, disableNotification *bool, replyToMessageID *int32, replyMarkup *interface{}) (*Message, error) {

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
			name:     "sticker",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: sticker,
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

	res, err := obj.callMethod("sendSticker", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*Message), nil
}

func (obj *BotAPI) GetStickerSet(name string) (*StickerSet, error) {

	parameters := []parameter{
		{
			name:     "name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: name,
		},
	}

	res, err := obj.callMethod("getStickerSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*StickerSet), nil
}

func (obj *BotAPI) UploadStickerFile(userID int32, pngSticker InputFile) (*File, error) {

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
			name:     "png_sticker",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
			},
			value: pngSticker,
		},
	}

	res, err := obj.callMethod("uploadStickerFile", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*File), nil
}

func (obj *BotAPI) CreateNewStickerSet(userID int32, name string, title string, pngSticker interface{}, emojis string, containsMasks *bool, maskPosition *MaskPosition) (*bool, error) {

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
			name:     "name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: name,
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
			name:     "png_sticker",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: pngSticker,
		},
		{
			name:     "emojis",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: emojis,
		},
		{
			name:     "contains_masks",
			required: false,
			types: []parameterType{
				parameterTypeBoolean,
			},
			value: containsMasks,
		},
		{
			name:     "mask_position",
			required: false,
			types: []parameterType{
				parameterTypeMaskPosition,
			},
			value: maskPosition,
		},
	}

	res, err := obj.callMethod("createNewStickerSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) AddStickerToSet(userID int32, name string, pngSticker interface{}, emojis string, maskPosition *MaskPosition) (*bool, error) {

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
			name:     "name",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: name,
		},
		{
			name:     "png_sticker",
			required: true,
			types: []parameterType{
				parameterTypeInputFile,
				parameterTypeString,
			},
			value: pngSticker,
		},
		{
			name:     "emojis",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: emojis,
		},
		{
			name:     "mask_position",
			required: false,
			types: []parameterType{
				parameterTypeMaskPosition,
			},
			value: maskPosition,
		},
	}

	res, err := obj.callMethod("addStickerToSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) SetStickerPositionInSet(sticker string, position int32) (*bool, error) {

	parameters := []parameter{
		{
			name:     "sticker",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: sticker,
		},
		{
			name:     "position",
			required: true,
			types: []parameterType{
				parameterTypeInteger,
			},
			value: position,
		},
	}

	res, err := obj.callMethod("setStickerPositionInSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

func (obj *BotAPI) DeleteStickerFromSet(sticker string) (*bool, error) {

	parameters := []parameter{
		{
			name:     "sticker",
			required: true,
			types: []parameterType{
				parameterTypeString,
			},
			value: sticker,
		},
	}

	res, err := obj.callMethod("deleteStickerFromSet", parameters...)
	if err != nil {
		return nil, err
	}

	return res.(*bool), nil
}

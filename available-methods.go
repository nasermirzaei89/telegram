package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

func (obj *API) getMe() (*User, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("getMe", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *User `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendMessage(chatID interface{}, text string, parseMode *string, disableWebPagePreview, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("text", text)
	if parseMode != nil {
		writer.WriteField("parse_mode", *parseMode)
	}
	if disableWebPagePreview != nil {
		writer.WriteField("disable_web_page_preview", fmt.Sprintf("%t", *disableWebPagePreview))
	}
	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}
	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendMessage", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) forwardMessage(chatID, fromChatID interface{}, disableNotification *bool, messageID int64) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch fromChatID.(type) {
	case string:
		writer.WriteField("from_chat_id", fromChatID.(string))
	case int64:
		writer.WriteField("from_chat_id", fmt.Sprintf("%d", fromChatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for from chat id", fromChatID)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	writer.WriteField("message_id", fmt.Sprintf("%d", messageID))

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("forwardMessage", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendPhoto(chatID interface{}, photo interface{}, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch photo.(type) {
	case string:
		err := writer.WriteField("photo", photo.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("photo", "photo")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, photo.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for photo", photo)
	}

	if caption != nil {
		writer.WriteField("caption", *caption)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendPhoto", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendAudio(chatID interface{}, audio interface{}, caption *string, duration *int64, performer *string, title *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch audio.(type) {
	case string:
		err := writer.WriteField("audio", audio.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("audio", "audio")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, audio.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for audio", audio)
	}

	if caption != nil {
		writer.WriteField("caption", *caption)
	}

	if duration != nil {
		writer.WriteField("duration", fmt.Sprintf("%d", *duration))
	}

	if performer != nil {
		writer.WriteField("performer", *performer)
	}

	if title != nil {
		writer.WriteField("title", *title)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("sendAudio", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendDocument(chatID interface{}, document interface{}, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch document.(type) {
	case string:
		err := writer.WriteField("document", document.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("document", "document")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, document.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for document", document)
	}

	if caption != nil {
		writer.WriteField("caption", *caption)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendDocument", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendSticker(chatID interface{}, sticker interface{}, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch sticker.(type) {
	case string:
		err := writer.WriteField("sticker", sticker.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("sticker", "sticker")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, sticker.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for sticker", sticker)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendSticker", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendVideo(chatID interface{}, video interface{}, duration, width, height *int64, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch video.(type) {
	case string:
		err := writer.WriteField("video", video.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("video", "video")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, video.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for video", video)
	}

	if duration != nil {
		writer.WriteField("duration", fmt.Sprintf("%d", *duration))
	}

	if width != nil {
		writer.WriteField("width", fmt.Sprintf("%d", *width))
	}

	if height != nil {
		writer.WriteField("height", fmt.Sprintf("%d", *height))
	}

	if caption != nil {
		writer.WriteField("caption", *caption)
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendVideo", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendVoice(chatID interface{}, voice interface{}, caption *string, duration *int64, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		err := writer.WriteField("chat_id", chatID.(string))
		if err != nil {
			return nil, err
		}
	case int64:
		err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	switch voice.(type) {
	case string:
		err := writer.WriteField("voice", voice.(string))
		if err != nil {
			return nil, err
		}
	case InputFile:
		part, err := writer.CreateFormFile("voice", "voice")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, voice.(InputFile))
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalide type %T for voice", voice)
	}

	if caption != nil {
		writer.WriteField("caption", *caption)
	}

	if duration != nil {
		writer.WriteField("duration", fmt.Sprintf("%d", *duration))
	}

	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}

	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}
	rsp, err := obj.makingRequest("sendVoice", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendLocation(chatID interface{}, latitude, longitude float64, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("latitude", fmt.Sprintf("%f", latitude))
	writer.WriteField("longitude", fmt.Sprintf("%f", longitude))
	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}
	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("sendLocation", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendVenue(chatID interface{}, latitude, longitude float64, title, address string, foursquareID *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("latitude", fmt.Sprintf("%f", latitude))
	writer.WriteField("longitude", fmt.Sprintf("%f", longitude))
	writer.WriteField("title", title)
	writer.WriteField("address", address)
	if foursquareID != nil {
		writer.WriteField("foursquare_id", *foursquareID)
	}
	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}
	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("sendVenue", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) sendContact(chatID interface{}, phoneNumber string, firstName string, lastName *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("phone_number", phoneNumber)
	writer.WriteField("first_name", firstName)
	if lastName != nil {
		writer.WriteField("last_name", *lastName)
	}
	if disableNotification != nil {
		writer.WriteField("disable_notification", fmt.Sprintf("%t", *disableNotification))
	}
	if replyToMessageID != nil {
		writer.WriteField("reply_to_message_id", fmt.Sprintf("%d", *replyToMessageID))
	}
	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("sendContact", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

// Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data.
const (
	ChatActionTyping         = "typing"
	ChatActionUploadPhoto    = "upload_photo"
	ChatActionRecordVideo    = "record_video"
	ChatActionUploadVideo    = "upload_video"
	ChatActionRecordAudio    = "record_audio"
	ChatActionUploadAudio    = "upload_audio"
	ChatActionUploadDocument = "upload_document"
	ChatActionFindLocation   = "find_location"
)

func (obj *API) sendChatAction(chatID interface{}, action string) (*bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("action", action)

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("sendChatAction", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *bool `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getUserProfilePhotos(userID int64, offset, limit *int64) (*UserProfilePhotos, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("user_id", fmt.Sprintf("%d", userID))
	if offset != nil {
		writer.WriteField("offset", fmt.Sprintf("%d", *offset))
	}
	if limit != nil {
		writer.WriteField("limit", fmt.Sprintf("%d", *limit))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getUserProfilePhotos", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *UserProfilePhotos `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getFile(fileID string) (*File, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("file_id", fileID)

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getFile", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *File `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) kickChatMember(chatID interface{}, userID int64) (*bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("user_id", fmt.Sprintf("%d", userID))

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("kickChatMember", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *bool `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) leaveChat(chatID interface{}) (*bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("leaveChat", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *bool `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) unbanChatMember(chatID interface{}, userID int64) (*bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("user_id", fmt.Sprintf("%d", userID))

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("unbanChatMember", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *bool `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getChat(chatID interface{}) (*Chat, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getChat", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *Chat `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getChatAdministrators(chatID interface{}) (*[]ChatMember, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getChatAdministrators", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *[]ChatMember `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getChatMembersCount(chatID interface{}) (*int64, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getChatMembersCount", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *int64 `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) getChatMember(chatID interface{}, userID int64) (*ChatMember, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	writer.WriteField("user_id", fmt.Sprintf("%d", userID))

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getChatMember", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *ChatMember `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

func (obj *API) answerCallbackQuery(callbackQueryID string, text *string, showAlert *bool, u *string, cacheTime *int64) (*bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("callback_query_id", callbackQueryID)

	if text != nil {
		writer.WriteField("text", *text)
	}

	if showAlert != nil {
		writer.WriteField("show_alert", fmt.Sprintf("%t", *showAlert))
	}

	if u != nil {
		writer.WriteField("url", *u)
	}

	if cacheTime != nil {
		writer.WriteField("cache_time", fmt.Sprintf("%d", *cacheTime))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("answerCallbackQuery", writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result *bool `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, err
	}

	if !v.OK {
		return nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil
}

package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
)

func (obj *API) editMessageText(chatID interface{}, messageID *int64, inlineMessageID *string, text string, parseMode *string, disableWebPagePreview *bool, replyMarkup interface{}) (*Message, *bool, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	switch chatID.(type) {
	case string:
		writer.WriteField("chat_id", chatID.(string))
	case int64:
		writer.WriteField("chat_id", fmt.Sprintf("%d", chatID.(int64)))
	default:
		return nil, nil, fmt.Errorf("invalide type %T for chat id", chatID)
	}

	if messageID != nil {
		writer.WriteField("message_id", fmt.Sprintf("%d", *messageID))
	}

	if inlineMessageID != nil {
		writer.WriteField("inline_message_id", *inlineMessageID)
	}

	writer.WriteField("text", text)

	if parseMode != nil {
		writer.WriteField("parse_mode", *parseMode)
	}

	if disableWebPagePreview != nil {
		writer.WriteField("disable_web_page_preview", fmt.Sprintf("%t", *disableWebPagePreview))
	}

	if replyMarkup != nil {
		b, err := json.Marshal(replyMarkup)
		if err != nil {
			return nil, nil, err
		}
		writer.WriteField("reply_markup", string(b))
	}

	err := writer.Close()
	if err != nil {
		return nil, nil, err
	}

	rsp, err := obj.makingRequest("editMessageText", writer.FormDataContentType(), body)
	if err != nil {
		return nil, nil, err
	}

	v := new(struct {
		Result *Message `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, v)
	if err != nil {
		return nil, nil, err
	}

	if !v.OK {
		return nil, nil, fmt.Errorf("error %d: %s", v.ErrorCode, v.Description)
	}

	return v.Result, nil, nil
}

func (obj *API) editMessageCaption(chatID interface{}, messageID int64, inlineMessageID int64, caption string, replyMarkup interface{}) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

func (obj *API) editMessageReplyMarkup(chatID interface{}, messageID int64, inlineMessageID int64, replyMarkup interface{}) (*Message, *bool, error) {
	return nil, nil, errors.New("Not implemented")
}

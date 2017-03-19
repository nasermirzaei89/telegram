package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"
)

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
}

// List the types of updates you want your bot to receive.
// For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
// See Update for a complete list of available update types.
// Specify an empty list to receive all updates regardless of type (default).
// If not specified, the previous setting will be used.
// Please note that this parameter doesn't affect updates created before the call to the getUpdates,
// so unwanted updates may be received for a short period of time.
const (
	UpdateMessage            = "message"
	UpdateEditedMessage      = "edited_message"
	UpdateChannelPost        = "channel_post"
	UpdateEditedChannelPost  = "edited_channel_post"
	UpdateInlineQuery        = "inline_query"
	UpdateChosenInlineResult = "chosen_inline_result"
	UpdateCallbackQuery      = "callback_query"
)

func (obj *API) getUpdates(offset, limit, timeout int64, allowedUpdates []string) ([]Update, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	allowedUpdatesStr := "["
	for i := range allowedUpdates {
		allowedUpdatesStr += fmt.Sprintf(`"%s",`, allowedUpdates[i])
	}
	allowedUpdatesStr = strings.TrimRight(allowedUpdatesStr, ",") + "]"

	writer.WriteField("offset", fmt.Sprintf("%d", offset))
	writer.WriteField("limit", fmt.Sprintf("%d", limit))
	writer.WriteField("timeout", fmt.Sprintf("%d", timeout))
	writer.WriteField("allowed_updates", allowedUpdatesStr)

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest("getUpdates", writer.FormDataContentType(), body)

	if err != nil {
		return nil, err
	}

	v := new(struct {
		Result []Update `json:"result,omitempty"`
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

// setWebhook
// TODO: implement this

// deleteWebhook
// TODO: implement this

// getWebhookInfo
// TODO: implement this

// WebhookInfo
// TODO: implement this

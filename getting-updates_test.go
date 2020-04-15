package telegram_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasermirzaei89/telegram"
	"github.com/stretchr/testify/assert"
)

func TestGetUpdates(t *testing.T) {
	response := []byte(`{
  "ok": true,
  "result": [
    {
      "update_id": 109399605,
      "message": {
        "message_id": 1234,
        "from": {
          "id": 123456789,
          "is_bot": false,
          "first_name": "John",
          "last_name": "Doe",
          "username": "johndoe",
          "language_code": "en"
        },
        "chat": {
          "id": 123456789,
          "first_name": "John",
          "last_name": "Doe",
          "username": "johndoe",
          "type": "private"
        },
        "date": 1234567890,
        "text": "Test"
      }
    }
  ]
}`)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := fmt.Sprintf("/bot%s/getUpdates", testToken)
		if r.URL.String() != u {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"ok":false,"error_code":401,"description":"Unauthorized"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}))

	defer server.Close()

	// success
	bot := telegram.New(testToken, telegram.SetBaseURL(server.URL))

	res, err := bot.GetUpdates()
	assert.Nil(t, err)

	assert.True(t, res.IsOK())
	assert.Zero(t, res.GetErrorCode())
	assert.NotNil(t, res.GetUpdates())

	for _, u := range res.GetUpdates() {
		assert.NotZero(t, u.UpdateID)
	}

	// fail
	bot = telegram.New(invalidToken, telegram.SetBaseURL(server.URL))

	res, err = bot.GetUpdates()
	assert.Nil(t, err)

	assert.False(t, res.IsOK())
	assert.Empty(t, len(res.GetUpdates()))
	assert.Equal(t, http.StatusUnauthorized, res.GetErrorCode())
	assert.Equal(t, "Unauthorized", res.GetDescription())
	assert.Nil(t, res.GetParameters())
}

func TestSetWebhook(t *testing.T) {
	response := []byte(`{"ok":true,"result":true,"description":"Webhook was set"}`)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := fmt.Sprintf("/bot%s/setWebhook", testToken)
		if r.URL.String() != u {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"ok":false,"error_code":401,"description":"Unauthorized"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}))

	defer server.Close()

	// success
	bot := telegram.New(testToken, telegram.SetBaseURL(server.URL))

	res, err := bot.SetWebhook(
		telegram.SetURL("https://example.com/telegram/webhook"),
		telegram.SetMaxConnections(40),
		telegram.SetAllowedUpdates("message"),
	)
	assert.Nil(t, err)

	assert.True(t, res.IsOK())
	assert.Zero(t, res.GetErrorCode())
	assert.Equal(t, "Webhook was set", res.GetDescription())
}

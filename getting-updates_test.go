package telegram_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasermirzaei89/telegram"
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
	if err != nil {
		t.Errorf("error on get updates: %s", err.Error())
		return
	}

	if !res.IsOK() {
		t.Error("result should be ok but is not")
		return
	}

	if res.GetErrorCode() != 0 {
		t.Errorf("result error code should be zero but is %d", res.GetErrorCode())
		return
	}

	updates := res.GetUpdates()

	if updates == nil {
		t.Error("result updates should be array but are nil")
		return
	}

	for _, u := range updates {
		if u.UpdateID == 0 {
			t.Error("update is should not be zero but is")
		}
	}

	// fail
	bot = telegram.New(invalidToken, telegram.SetBaseURL(server.URL))

	res, err = bot.GetUpdates()
	if err != nil {
		t.Errorf("error on get updates: %s", err.Error())
		return
	}

	if res.IsOK() {
		t.Error("result should not be ok but is")
		return
	}

	if res.GetErrorCode() != http.StatusUnauthorized {
		t.Errorf("result error code should be %d but is %d", http.StatusUnauthorized, res.GetErrorCode())
		return
	}
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
	if err != nil {
		t.Errorf("error on set web hook: %s", err.Error())
		return
	}

	if !res.IsOK() {
		t.Error("result should be ok but is not")
		return
	}

	if res.GetErrorCode() != 0 {
		t.Errorf("result error code should be zero but is %d", res.GetErrorCode())
		return
	}

	expectedDesc := "Webhook was set"
	desc := res.GetDescription()
	if desc != expectedDesc {
		t.Errorf("result description should be '%s' but is '%s'", expectedDesc, desc)
	}
}

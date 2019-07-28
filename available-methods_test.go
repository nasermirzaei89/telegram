package telegram_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasermirzaei89/telegram"
	"github.com/stretchr/testify/assert"
)

const (
	testToken    = "someToken"
	invalidToken = "invalidToken"
)

func TestGetMe(t *testing.T) {
	response := []byte(`{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"Test Bot","username":"TestBot"}}`)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := fmt.Sprintf("/bot%s/getMe", testToken)
		if r.URL.String() != u {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"ok":false,"error_code":401,"description":"Unauthorized"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}))

	defer server.Close()

	telegram.BaseURL = server.URL

	// success
	bot := telegram.New(testToken)

	res, err := bot.GetMe()
	assert.Nil(t, err)

	assert.True(t, res.IsOK())
	assert.Zero(t, res.GetErrorCode())
	assert.NotNil(t, res.GetUser())
	assert.Equal(t, 1, res.GetUser().ID)
	assert.True(t, res.GetUser().IsBot)
	assert.Equal(t, "Test Bot", res.GetUser().FirstName)
	assert.Nil(t, res.GetUser().LastName)
	assert.NotNil(t, res.GetUser().Username)
	assert.Equal(t, "TestBot", *res.GetUser().Username)
	assert.Nil(t, res.GetUser().LanguageCode)

	// fail
	bot = telegram.New(invalidToken)

	res, err = bot.GetMe()
	assert.Nil(t, err)

	assert.False(t, res.IsOK())
	assert.Nil(t, res.GetUser())
	assert.Equal(t, http.StatusUnauthorized, res.GetErrorCode())
	assert.Equal(t, "Unauthorized", res.GetDescription())
	assert.Nil(t, res.GetParameters())
}

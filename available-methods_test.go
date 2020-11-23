package telegram_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasermirzaei89/telegram"
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

	// success
	bot := telegram.New(testToken, telegram.SetBaseURL(server.URL))

	res, err := bot.GetMe()
	if err != nil {
		t.Errorf("error on get me: %s", err.Error())
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

	usr := res.GetUser()

	if usr == nil {
		t.Error("result user should not be nil but is")
		return
	}

	if usr.ID != 1 {
		t.Errorf("result user id should be 1 but is %d", usr.ID)
	}

	if !usr.IsBot {
		t.Errorf("result user should be bot but is not")
	}

	expectedFirstName := "Test Bot"
	if usr.FirstName != expectedFirstName {
		t.Errorf("result user first name should be '%s' but is '%s'", expectedFirstName, usr.FirstName)
	}

	if usr.LastName != nil {
		t.Errorf("result user should not have last name but has")
	}

	if usr.Username == nil {
		t.Errorf("result user username should not be nil but is")
		return
	}

	expectedUsername := "TestBot"
	if *usr.Username != expectedUsername {
		t.Errorf("result user username should be '%s' but is '%s'", expectedUsername, *usr.Username)
	}

	// fail
	bot = telegram.New(invalidToken, telegram.SetBaseURL(server.URL))

	res, err = bot.GetMe()
	if err != nil {
		t.Errorf("error on get me: %s", err.Error())
		return
	}

	if res.IsOK() {
		t.Error("result should not be ok but is")
		return
	}

	if res.GetUser() != nil {
		t.Error("result user should be nil but is not")
	}

	if res.GetErrorCode() != http.StatusUnauthorized {
		t.Errorf("result error code should be %d but is %d", http.StatusUnauthorized, res.GetErrorCode())
		return
	}
}

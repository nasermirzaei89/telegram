package telegram

import "net/http"

// API is telegram bot api structure
type API struct {
	token  string
	offset int64
	client *http.Client
}

// NewAPI will return new telegram bot api
func NewAPI(token string) *API {
	api := new(API)
	api.token = token
	api.offset = -1
	api.client = new(http.Client)
	return api
}

package telegram

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Error is structure of error embedded in telegram responses
type Error struct {
	OK          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (obj *API) makingRequest(methodName string, contentType string, body *bytes.Buffer) ([]byte, error) {
	urlStr := fmt.Sprintf("https://api.telegram.org/bot%s/%s", obj.token, methodName)

	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	rsp, err := obj.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		er := rsp.Body.Close()
		if er != nil {
			log.Printf("Error on close io: %s", er)
		}
	}()

	return ioutil.ReadAll(rsp.Body)
}

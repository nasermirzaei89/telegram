package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Response struct {
	OK          bool                `json:"ok"`
	Description *string             `json:"description,omitempty"`
	ErrorCode   *int                `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      interface{}         `json:"result,omitempty"`
}

type request struct {
	url    string
	body   *bytes.Buffer
	writer *multipart.Writer
	err    error
}

func newRequest(token, methodName string) *request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	return &request{
		url:    fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methodName),
		body:   body,
		writer: writer,
	}
}

func (r *request) do(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	err := r.writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, r.url, r.body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}

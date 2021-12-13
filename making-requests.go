package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// Response general interface.
type Response interface {
	IsOK() bool
	GetErrorCode() int
	GetDescription() string
	GetParameters() *ResponseParameters
}

type response struct {
	OK          bool                `json:"ok"`
	Description *string             `json:"description,omitempty"`
	ErrorCode   *int                `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (r *response) IsOK() bool {
	return r.OK
}

func (r *response) GetErrorCode() int {
	if r.ErrorCode != nil {
		return *r.ErrorCode
	}

	return 0
}

func (r *response) GetDescription() string {
	if r.Description != nil {
		return *r.Description
	}

	return ""
}

func (r *response) GetParameters() *ResponseParameters {
	return r.Parameters
}

type request struct {
	params map[string]interface{}
}

func (b *Bot) doRequest(ctx context.Context, methodName string, res interface{}, options ...MethodOption) error {
	r := request{
		params: map[string]interface{}{},
	}

	for i := range options {
		options[i](&r)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range r.params {
		switch t := v.(type) {
		case io.Reader:
			w, err := writer.CreateFormFile(k, k)
			if err != nil {
				return fmt.Errorf("error on create field from file: %w", err)
			}

			b, err := ioutil.ReadAll(t)
			if err != nil {
				return fmt.Errorf("error on read: %w", err)
			}

			_, err = w.Write(b)
			if err != nil {
				return fmt.Errorf("error on write data: %w", err)
			}
		case string:
			err := writer.WriteField(k, t)
			if err != nil {
				return fmt.Errorf("error on write field: %w", err)
			}
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return fmt.Errorf("error on marshal json: %w", err)
			}

			err = writer.WriteField(k, string(b))
			if err != nil {
				return fmt.Errorf("error on write field: %w", err)
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return fmt.Errorf("error on close writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/bot%s/%s", b.baseURL, b.token, methodName), body)
	if err != nil {
		return fmt.Errorf("error on new http request: %w", err)
	}

	// don't send content type on empty body
	if len(r.params) > 0 {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error on do http request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return fmt.Errorf("error decode json response: %w", err)
	}

	return nil
}

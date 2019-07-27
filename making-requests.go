package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

// BaseURL is base of telegram api url
var BaseURL = "https://api.telegram.org"

// Response general interface
type Response interface {
	IsOK() bool
	GetErrorCode() *int
	GetDescription() *string
	GetParameters() *ResponseParameters
}

type response struct {
	OK          bool                `json:"ok"`
	Description *string             `json:"description,omitempty"`
	ErrorCode   *int                `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      interface{}         `json:"result,omitempty"`
}

func (r *response) IsOK() bool {
	return r.OK
}

func (r *response) GetErrorCode() *int {
	return r.ErrorCode
}

func (r *response) GetDescription() *string {
	return r.Description
}

func (r *response) GetParameters() *ResponseParameters {
	return r.Parameters
}

type request struct {
	url    string
	body   *bytes.Buffer
	writer *multipart.Writer
	err    error
}

func doRequest(token, methodName string, res interface{}, options ...Option) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	req := request{
		url:    fmt.Sprintf("%s/bot%s/%s", BaseURL, token, methodName),
		body:   body,
		writer: writer,
	}

	for i := range options {
		options[i](&req)
	}

	if req.err != nil {
		return req.err
	}

	return (&req).do(&res)
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

	if r.body.Len() > 68 {
		req.Header.Set("Content-Type", r.writer.FormDataContentType())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

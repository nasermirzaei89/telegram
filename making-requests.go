package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	params map[string]interface{}
}

func doRequest(token, methodName string, res interface{}, options ...Option) error {
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
				return err
			}

			b, err := ioutil.ReadAll(t)
			if err != nil {
				return err
			}

			_, err = w.Write(b)
			if err != nil {
				return err
			}
		case string:
			err := writer.WriteField(k, t)
			if err != nil {
				return err
			}
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return err
			}

			err = writer.WriteField(k, string(b))
			if err != nil {
				return err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/bot%s/%s", BaseURL, token, methodName), body)
	if err != nil {
		return err
	}

	// don't send content type on empty body
	if len(r.params) > 0 {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return err
	}

	return nil
}

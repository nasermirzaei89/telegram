package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Response interface {
	IsOK() bool
	Error() ResponseError
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

func (r *response) Error() ResponseError {
	if !r.IsOK() {
		return &responseError{
			description: r.Description,
			errorCode:   r.ErrorCode,
			parameters:  r.Parameters,
		}
	}

	return nil
}

type ResponseError interface {
	GetDescription() string
	GetErrorCode() int
	GetParameters() *ResponseParameters
}

type responseError struct {
	description *string
	errorCode   *int
	parameters  *ResponseParameters
}

func (r *responseError) GetDescription() string {
	if r.description != nil {

		return *r.description
	}

	return ""
}

func (r *responseError) GetErrorCode() int {
	if r.errorCode != nil {
		return *r.errorCode
	}

	return 0
}

func (r *responseError) GetParameters() *ResponseParameters {
	return r.parameters
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
		url:    fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methodName),
		body:   body,
		writer: writer,
	}

	for i := range options {
		options[i](&req)
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

	req.Header.Set("Content-Type", r.writer.FormDataContentType())

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

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

// Error is structure of error embedded in telegram responses
type Error struct {
	OK          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

// Error ...
func (e Error) Error() string {
	if !e.OK {
		return fmt.Sprintf("error %d: %s", e.ErrorCode, e.Description)
	}

	return ""
}

type parameterType int

const (
	parameterTypeString = iota
	parameterTypeInteger
	parameterTypeFloatNumber
	parameterTypeBoolean
	parameterTypeInputFile
	parameterTypeInlineKeyboardMarkup
	parameterTypeReplyKeyboardMarkup
	parameterTypeReplyKeyboardRemove
	parameterTypeForceReply

	parameterTypeArrayOfInputMedia
	parameterTypeArrayOfString
	parameterTypeMaskPosition
	parameterTypeArrayOfLabeledPrice
	parameterTypeArrayOfShippingOption
	parameterTypeArrayOfInlineQueryResult
)

type parameter struct {
	name     string
	required bool
	value    interface{}
	types    []parameterType
}

func (obj *BotAPI) callMethod(name string, parameters ...parameter) (interface{}, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// TODO: Check parameter types
	for i := range parameters {
		if !parameters[i].required && parameters[i].value == nil {
			continue
		}

		writeField(writer, parameters[i].name, parameters[i].value)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	rsp, err := obj.makingRequest(name, writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	res := new(struct {
		Result interface{} `json:"result,omitempty"`
		Error
	})

	err = json.Unmarshal(rsp, res)
	if err != nil {
		return nil, err
	}

	if !res.OK {
		return nil, Error{OK: res.OK, ErrorCode: res.ErrorCode, Description: res.Description}
	}

	return res.Result, nil
}

func writeField(writer *multipart.Writer, name string, value interface{}) error {
	switch v := value.(type) {
	case string:
		err := writer.WriteField(name, v)
		if err != nil {
			return err
		}
	case int32, int64:
		err := writer.WriteField(name, fmt.Sprintf("%d", v))
		if err != nil {
			return err
		}
	case InputFile:
		part, err := writer.CreateFormFile(name, name)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, v)
		if err != nil {
			return err
		}
	case bool:
		err := writer.WriteField(name, fmt.Sprintf("%t", v))
		if err != nil {
			return err
		}
	case InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply:
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}

		err = writer.WriteField(name, string(b))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalide type %T for %s", name, v)
	}

	return nil
}

func (obj *BotAPI) makingRequest(methodName string, contentType string, body *bytes.Buffer) ([]byte, error) {
	urlStr := fmt.Sprintf("https://api.telegram.org/bot%s/%s", obj.Token, methodName)

	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	cl := new(http.Client)

	rsp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}

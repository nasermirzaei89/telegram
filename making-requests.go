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

// request content types
const (
	// URL query string // we doesn't support this
	ContentTypeApplicationXWWWFormURLEncode = "application/x-www-form-urlencoded"
	ContentTypeApplicationJSON              = "application/json"    // except for uploading files
	ContentTypeMultipartFormData            = "multipart/form-data" // use to upload files
)

// Error is structure of error embedded in telegram responses
type Error struct {
	OK          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

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

		switch v := parameters[i].value.(type) {
		case string:
			err := writer.WriteField(parameters[i].name, v)
			if err != nil {
				return nil, err
			}
		case int32, int64:
			err := writer.WriteField(parameters[i].name, fmt.Sprintf("%d", v))
			if err != nil {
				return nil, err
			}
		case InputFile:
			part, err := writer.CreateFormFile(parameters[i].name, parameters[i].name)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, v)
			if err != nil {
				return nil, err
			}
		case bool:
			err := writer.WriteField(parameters[i].name, fmt.Sprintf("%t", v))
			if err != nil {
				return nil, err
			}
		case InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply:
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}

			err = writer.WriteField(parameters[i].name, string(b))
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalide type %T for %s", parameters[i].name, v)
		}
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

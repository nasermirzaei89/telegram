package telegram

import "fmt"

type Response struct {
	OK          bool                `json:"ok"`
	Description *string             `json:"description,omitempty"`
	ErrorCode   *int                `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      interface{}         `json:"result,omitempty"`
}

func (b *bot) getURL(methodName string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, methodName)
}

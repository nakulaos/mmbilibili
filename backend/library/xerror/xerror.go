package xerror

import (
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type XError struct {
	Code             int32             `json:"code"`
	Message          i18n.Message      `json:"-"`
	TemplateData     map[string]string `json:"-"`
	TranslateMessage string            `json:"message"`
	Status           *status.Status
}

func (e *XError) BizStatusCode() int32 {
	return e.Code
}

func (e *XError) BizMessage() string {
	return e.Message.ID
}

func (e *XError) BizExtra() map[string]string {
	return e.TemplateData
}

func (e *XError) Error() string {
	data, err := json.Marshal(e)
	if err != nil || e.TranslateMessage == "" {
		return e.Message.ID
	}
	return string(data)
}

func (e *XError) SetTranslateMessage(m string) *XError {
	err := NewXError(e.Code, e.Message)
	err.TranslateMessage = m
	return err
}

func (e *XError) WithTemplateData(data map[string]string) *XError {
	err := NewXError(e.Code, e.Message)
	err.TemplateData = data
	return err
}

func NewXError(code int32, message i18n.Message) *XError {
	return &XError{Code: code, Message: message}
}

/**
 ******************************************************************************
 * @file           : errorcode.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package xerror

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type XError struct {
	Code             int                    `json:"code"`
	Message          i18n.Message           `json:"-"`
	TemplateData     map[string]interface{} `json:"-"`
	TranslateMessage string                 `json:"message"`
}

func (e XError) Error() string {
	data, err := json.Marshal(e)
	if err != nil || e.TranslateMessage == "" {
		return e.Message.ID
	}
	return string(data)
}

func NewXError(code int, message i18n.Message) XError {
	return XError{Code: code, Message: message}
}

func (e XError) WithTemplateData(data map[string]interface{}) XError {
	e.TemplateData = data
	return e
}

func (e XError) ErrorCode() int {
	return e.Code
}

func (e XError) SetTranslateMessage(m string) XError {
	e.TranslateMessage = m
	return e
}

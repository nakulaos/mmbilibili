/**
 ******************************************************************************
 * @file           : common.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/25
 ******************************************************************************
 */

package xerror

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	Ok XError = NewXError(0, i18n.Message{
		ID:    "Ok",
		Other: "Success",
	})

	AuthorizationError XError = NewXError(100001, i18n.Message{
		ID:    "AuthorizationError",
		Other: "Authorization error",
	})

	ServerError XError = NewXError(100002, i18n.Message{
		ID:    "ServerError",
		Other: "Server error",
	})

	InvalidParamsError XError = NewXError(100003, i18n.Message{
		ID:    "InvalidParamsError",
		Other: "Invalid parameters {{.Params}}",
	})
)

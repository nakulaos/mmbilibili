package ecode

import (
	"backend/library/xerror"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	Ok *xerror.XError = xerror.NewXError(0, i18n.Message{
		ID:    "Ok",
		Other: "Success",
	})

	AuthorizationError *xerror.XError = xerror.NewXError(100001, i18n.Message{
		ID:    "AuthorizationError",
		Other: "Authorization error",
	})

	ServerError *xerror.XError = xerror.NewXError(100002, i18n.Message{
		ID:    "ServerError",
		Other: "Server error",
	})

	InvalidParamsError *xerror.XError = xerror.NewXError(100003, i18n.Message{
		ID:    "InvalidParamsError",
		Other: "Invalid parameters {{.Params}}",
	})

	RefreshTokenError *xerror.XError = xerror.NewXError(100004, i18n.Message{
		ID:    "RefreshTokenError",
		Other: "Refresh token error",
	})
)

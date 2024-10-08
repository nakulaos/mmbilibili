package xerror

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	CommentNotExistError XError = NewXError(400001, i18n.Message{
		ID:    "CommentNotExistError",
		Other: "The comment is not exist",
	})
)

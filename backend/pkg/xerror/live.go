package xerror

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	LiveNotExistError XError = NewXError(300001, i18n.Message{
		ID:    "LiveNotExistError",
		Other: "The live is not exist",
	})

	LiveIsOverError XError = NewXError(300002, i18n.Message{
		ID:    "LiveIsOverError",
		Other: "The live is over",
	})

	LiveLikeIsExistError XError = NewXError(300003, i18n.Message{
		ID:    "LiveLikeIsExistError",
		Other: "The live like is exist",
	})

	LiveLikeNotExistError XError = NewXError(300004, i18n.Message{
		ID:    "LiveLikeNotExistError",
		Other: "The live like is not exist",
	})
)

package ecode

import (
	"backend/library/xerror"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	FileMaxSizeError *xerror.XError = xerror.NewXError(400001, i18n.Message{
		ID:    "FileMaxSizeError",
		Other: "The file size exceeds the limit",
	})

	FileNoAuthError *xerror.XError = xerror.NewXError(400002, i18n.Message{
		ID:    "FileNoAuthError",
		Other: "No permission to access the file",
	})

	FileNotExistError *xerror.XError = xerror.NewXError(400003, i18n.Message{
		ID:    "FileNotExistError",
		Other: "The file does not exist",
	})

	FileUploadError *xerror.XError = xerror.NewXError(400004, i18n.Message{
		ID:    "FileUploadError",
		Other: "Error uploading the file",
	})

	FileSizeIllegalError *xerror.XError = xerror.NewXError(400005, i18n.Message{
		ID:    "FileSizeIllegalError",
		Other: "The file size is illegal",
	})

	FileChunkSizeIllegalError *xerror.XError = xerror.NewXError(400006, i18n.Message{
		ID:    "FileChunkSizeIllegalError",
		Other: "The file chunk size is illegal",
	})
)

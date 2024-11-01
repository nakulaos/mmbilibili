package ecode

import (
	"backend/library/xerror"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	// UserNotExistError 用户不存在
	UserNotExistError *xerror.XError = xerror.NewXError(200001, i18n.Message{
		ID:    "UserNotExistError",
		Other: "The {{.UID}} is not exist",
	})
	// UserExistError 用户已存在
	UserExistError *xerror.XError = xerror.NewXError(200002, i18n.Message{
		ID:    "UserExistError",
		Other: "The {{.UID}} is exist",
	})
	// UserPasswordError 用户密码错误
	UserPasswordError *xerror.XError = xerror.NewXError(200003, i18n.Message{
		ID:    "UserPasswordError",
		Other: "The password is incorrect",
	})
	// PhoneExistError 手机号已存在
	PhoneExistError *xerror.XError = xerror.NewXError(200005, i18n.Message{
		ID:    "PhoneExistError",
		Other: "The phone number is already in use",
	})
	// EmailExistError 邮箱已存在
	EmailExistError *xerror.XError = xerror.NewXError(200006, i18n.Message{
		ID:    "EmailExistError",
		Other: "The email is already in use",
	})
	// PassWordError 密码错误
	PassWordError *xerror.XError = xerror.NewXError(200007, i18n.Message{
		ID:    "PassWordError",
		Other: "The password is incorrect",
	})
	// UserUploadVideoError 用户上传视频错误
	UserUploadVideoError *xerror.XError = xerror.NewXError(200008, i18n.Message{
		ID:    "UserUploadVideoError",
		Other: "Error uploading the video",
	})

	UserDisableError *xerror.XError = xerror.NewXError(200009, i18n.Message{
		ID:    "UserDisableError",
		Other: "The http is disabled",
	})

	UserNotAllowLiveError *xerror.XError = xerror.NewXError(200010, i18n.Message{
		ID:    "UserNotAllowLiveError",
		Other: "The http is not allowed to live",
	})

	FollowingIsExistError *xerror.XError = xerror.NewXError(200011, i18n.Message{
		ID:    "FollowingIsExistError",
		Other: "The following is exist",
	})

	FollowingIsNotExistError *xerror.XError = xerror.NewXError(200012, i18n.Message{
		ID:    "FollowingIsNotExistError",
		Other: "The following is not exist",
	})

	UserRoleNotAllowError *xerror.XError = xerror.NewXError(200013, i18n.Message{
		ID:    "UserRoleNotAllowError",
		Other: "The http role right is not allowed",
	})

	// relation

	UserNotFollowingSelfError *xerror.XError = xerror.NewXError(200014, i18n.Message{
		ID:    "UserNotFollowingSelfError",
		Other: "The user can't following self",
	})

	UserBlackUserError *xerror.XError = xerror.NewXError(200015, i18n.Message{
		ID:    "UserBlackUserError",
		Other: "The user is black user",
	})

	UserAlreadyFollowingError *xerror.XError = xerror.NewXError(200016, i18n.Message{
		ID:    "UserAlreadyFollowingError",
		Other: "The user is already following",
	})

	UserAlreadyBlackError *xerror.XError = xerror.NewXError(200017, i18n.Message{
		ID:    "UserAlreadyBlackError",
		Other: "The user is already black",
	})

	UserAlreadyFriendError *xerror.XError = xerror.NewXError(200018, i18n.Message{
		ID:    "UserAlreadyFriendError",
		Other: "The user is already friend",
	})

	UserAlreadyWhisperError *xerror.XError = xerror.NewXError(200019, i18n.Message{
		ID:    "UserAlreadyWhisperError",
		Other: "The user is already whisper",
	})

	UserAlreadyNotFollowingError *xerror.XError = xerror.NewXError(200020, i18n.Message{
		ID:    "UserAlreadyNotFollowingError",
		Other: "The user is already not following",
	})

	UserAlreadyNotBlackError *xerror.XError = xerror.NewXError(200021, i18n.Message{
		ID:    "UserAlreadyNotBlackError",
		Other: "The user is already not black",
	})

	UserAlreadyNotFriendError *xerror.XError = xerror.NewXError(200022, i18n.Message{
		ID:    "UserAlreadyNotFriendError",
		Other: "The user is already not friend",
	})

	UserAlreadyNotWhisperError *xerror.XError = xerror.NewXError(200023, i18n.Message{
		ID:    "UserAlreadyNotWhisperError",
		Other: "The user is already not whisper",
	})

	UserFollowingMaxError *xerror.XError = xerror.NewXError(200024, i18n.Message{
		ID:    "UserFollowingMaxError",
		Other: "The user following max",
	})
)

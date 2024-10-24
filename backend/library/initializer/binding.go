package initializer

import (
	"backend/app/common/ecode"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"regexp"
	"unicode"
)

func InitValidator() *binding.ValidateConfig {
	validateConfig := &binding.ValidateConfig{}

	// 校验用户名
	validateConfig.MustRegValidateFunc("username", func(args ...interface{}) error {
		if len(args) != 1 {
			return fmt.Errorf("the args must be one")
		}
		username, ok := args[0].(string)
		if !ok {
			return fmt.Errorf("the args must be a string")
		}

		// 用户名只能是字母和数字，且至少包含一个字母
		regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
		if !regex.MatchString(username) {
			return fmt.Errorf("username must contain only letters and numbers")
		}
		containsLetter := false
		for _, char := range username {
			if unicode.IsLetter(char) {
				containsLetter = true
				break
			}
		}
		if !containsLetter {
			return fmt.Errorf("username must contain at least one letter")
		}
		return nil
	})

	// 校验密码
	validateConfig.MustRegValidateFunc("password", func(args ...interface{}) error {
		if len(args) != 1 {
			return fmt.Errorf("the args must be one")
		}
		password, ok := args[0].(string)
		if !ok {
			return fmt.Errorf("the args must be a string")
		}

		var (
			hasUpper   bool
			hasLower   bool
			hasNumber  bool
			hasSpecial bool
		)

		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		// 至少包含三种字符类型
		charTypeCount := 0
		if hasUpper {
			charTypeCount++
		}
		if hasLower {
			charTypeCount++
		}
		if hasNumber {
			charTypeCount++
		}
		if hasSpecial {
			charTypeCount++
		}

		if charTypeCount < 3 {
			return errors.New("password must contain at least three types of: uppercase letter, lowercase letter, number, or special character")
		}

		return nil
	})

	validateConfig.SetValidatorErrorFactory(func(fieldSelector, msg string) error {
		return ecode.InvalidParamsError.WithTemplateData(map[string]string{
			"Params": fmt.Sprintf("%s %s", fieldSelector, msg),
		})
	})

	return validateConfig
}
